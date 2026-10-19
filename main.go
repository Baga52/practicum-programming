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
package main

import "fmt"

func main() {
	var number int = 322
	var hundreds int = number / 100
	var tens int = (number / 10) % 10
	var ones int = number % 10

	if hundreds == tens || tens == ones || hundreds == ones {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

// Третья задача
package main

import "fmt"

func main() {
	var a int = 11111

	if a < 10 {
	} else if a < 100 {
		a = a / 10
	} else if a < 1000 {
		a = a / 100
	} else if a < 10000 {
		a = a / 1000
	} else if a == 10000 {
		a = a / 10000
	} else {
		fmt.Println("Вы превысили, число не должно превышать 10000")
	}
	fmt.Println(a)
}

// Четвёртая задача
package main

import "fmt"

func main() {
	var n int = 123123

	var leftPart int = n / 1000
	var rightPart int = n % 1000

	var sumLeft int = (leftPart / 100) + ((leftPart / 10) % 10) + (leftPart % 10)
	var sumRight int = (rightPart / 100) + ((rightPart / 10) % 10) + (rightPart % 10)

	if sumLeft == sumRight {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}