package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	usdCourse = 83.93
	eurCourse = 93.83
)

func getValidAmount(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		input = strings.Replace(input, ",", ".", 1)

		if strings.TrimSpace(input) == "" {
			fmt.Println("Ошибка: Ввод не может быть пустым. Пожалуйста, введите число.")
			continue
		}

		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка: Некорректный тип данных. Пожалуйста, введите число (например, 100 или 1.5).")
			continue
		}

		if val < 0 {
			fmt.Println("Ошибка: Число не может быть отрицательным. Пожалуйста, введите неотрицательное число.")
			continue
		}

		return val
	}
}

func convertUsd(usdCount float64) {
	totalRub := usdCount * usdCourse
	fmt.Printf("За %.2f USD вы получите %.2f рублей.\n", usdCount, totalRub)
}

func convertEur(eurCount float64) {
	totalRub := eurCount * eurCourse
	fmt.Printf("За %.2f EUR вы получите %.2f рублей.\n", eurCount, totalRub)
}

func convertEurUsd(eurToUsdCount float64) {
	totalEurToUsd := (eurToUsdCount * eurCourse) / usdCourse
	fmt.Printf("За %.2f EUR вы получите %.2f долларов.\n", eurToUsdCount, totalEurToUsd)
}

func main() {
	for {
		inputUsd := getValidAmount("Введите кол-во долларов: ")
		convertUsd(inputUsd)

		inputEur := getValidAmount("Введите кол-во евро: ")
		convertEur(inputEur)

		inputEurToUsd := getValidAmount("Введите кол-во евро которые конвертировать в доллары: ")
		convertEurUsd(inputEurToUsd)

		for {
			fmt.Println("----------------------------------------")
			fmt.Print("Нужно ли конвертировать еще купюры? (да/нет или yes/no): ")
			var answer string
			fmt.Scanln(&answer)
			fmt.Println("----------------------------------------")

			answer = strings.ToLower(answer)

			if answer == "нет" || answer == "no" {
				fmt.Println("Программа завершена. До свидания!")
				return
			} else if answer == "да" || answer == "yes" {
				break
			} else {
				fmt.Println("Ошибка: Неправильно написано! Пожалуйста, введите да/нет или yes/no.")
			}
		}
	}
}
