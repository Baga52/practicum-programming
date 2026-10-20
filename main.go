package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const generatePage = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Генератор чисел</title>
</head>
<body>
    <h1>Генератор случайных чисел</h1>
    <button onclick="window.location.reload()">Сгенерировать новое число</button>
    <p>%s</p>
</body>
</html>
`

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Стрижечка под 0.6")
	})

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "админ панель")
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		randomNumber := rng.Intn(100)
		resultText := fmt.Sprintf("Сгенерировано число: %d", randomNumber)

		fmt.Fprintf(w, generatePage, resultText)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
