package main

import (
	"bufio"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
	"sync"
)

const filename = "bookmarks.txt"

var (
	bookmarksMu sync.RWMutex
	bookmarks   = make(map[string]string)
)

func loadBookmarks() error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 2)
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			url := strings.TrimSpace(parts[1])
			if name != "" && url != "" {
				bookmarks[name] = url
			}
		}
	}

	return scanner.Err()
}

func saveBookmarks() error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for name, url := range bookmarks {
		_, err := fmt.Fprintf(writer, "%s|%s\n", name, url)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func main() {
	if err := loadBookmarks(); err != nil {
		fmt.Printf("Ошибка при загрузке файла: %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string

		if r.Method == http.MethodPost {
			r.ParseForm()
			action := r.FormValue("action")
			name := strings.TrimSpace(r.FormValue("name"))
			url := strings.TrimSpace(r.FormValue("url"))

			if strings.Contains(name, "|") {
				message = "Ошибка: название не должно содержать символ '|'."
				action = "invalid" 
			}

			switch action {
			case "add":
				if name == "" || url == "" {
					message = "Ошибка: заполните все поля."
				} else {
					bookmarksMu.Lock()
					bookmarks[name] = url
					err := saveBookmarks()
					bookmarksMu.Unlock()

					if err != nil {
						message = "Ошибка при сохранении в TXT."
					} else {
						message = "Закладка успешно добавлена."
					}
				}
			case "delete":
				if name == "" {
					message = "Ошибка: укажите название закладки."
				} else {
					bookmarksMu.Lock()
					if _, exists := bookmarks[name]; exists {
						delete(bookmarks, name)
						err := saveBookmarks()
						if err != nil {
							message = "Ошибка при сохранении изменений."
						} else {
							message = "Закладка удалена."
						}
					} else {
						message = "Закладка не найдена."
					}
					bookmarksMu.Unlock()
				}
			}
		}

		searchQuery := strings.TrimSpace(r.URL.Query().Get("search"))
		searchLower := strings.ToLower(searchQuery)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintln(w, "<h1>Закладки</h1>")
		if message != "" {
			fmt.Fprintln(w, "<p><b>"+html.EscapeString(message)+"</b></p>")
		}

		fmt.Fprintf(w, `
			<form method="GET" action="/">
				<input name="search" placeholder="Поиск по названию или ссылке..." value="%s">
				<button type="submit">Найти</button>
				%s
			</form>
			<hr>
		`, html.EscapeString(searchQuery), func() string {
			if searchQuery != "" {
				return `<a href="/">Сбросить</a>`
			}
			return ""
		}())

		bookmarksMu.RLock()
		
		filteredBookmarks := make(map[string]string)
		for name, url := range bookmarks {
			if searchQuery == "" || 
			   strings.Contains(strings.ToLower(name), searchLower) || 
			   strings.Contains(strings.ToLower(url), searchLower) {
				filteredBookmarks[name] = url
			}
		}

		if len(bookmarks) == 0 {
			fmt.Fprintln(w, "<p>Список закладок пуст.</p>")
		} else if len(filteredBookmarks) == 0 {
			fmt.Fprintln(w, "<p>По вашему запросу ничего не найдено.</p>")
		} else {
			for name, url := range filteredBookmarks {
				safeName := html.EscapeString(name)
				safeURL := html.EscapeString(url)
				fmt.Fprintf(w, "<p>%s: <a href=\"%s\" target=\"_blank\">%s</a></p>", safeName, safeURL, safeURL)
			}
		}
		bookmarksMu.RUnlock()

		fmt.Fprintln(w, `
			<hr>
			<h3>Добавить закладку</h3>
			<form method="POST">
				<input type="hidden" name="action" value="add">
				<p><input name="name" placeholder="Название" required></p>
				<p><input name="url" placeholder="Ссылка" required></p>
				<p><button type="submit">Добавить</button></p>
			</form>
			
			<h3>Удалить закладку</h3>
			<form method="POST">
				<input type="hidden" name="action" value="delete">
				<p><input name="name" placeholder="Название" required></p>
				<p><button type="submit">Удалить</button></p>
			</form>
		`)
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
