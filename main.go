// Первая задача
package main

import "fmt"

func main() {
	var x int = 0

	if x > 0 {
		fmt.Println("Число положительное")
	} else if x < 0 {
		fmt.Println("Число отрицательное")
	} else {
		x = 0
		fmt.Println("Ноль")
	}
}

// Вторая задача
