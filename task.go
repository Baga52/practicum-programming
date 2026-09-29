package main

import "fmt"

// Задача 1.1

func main() {

	// - var с типом
	var x int = 1

	var f64 float64 = 34782147

	var password string = "password"

	var b bool = true

	var d byte = 123

	var r rune = 4

	fmt.Printf("Переменная: %T\n, %v\n", x, x)

	fmt.Printf("Переменная: %T\n, %v\n", f64, f64)

	fmt.Printf("Переменная: %T\n, %v\n", password, password)

	fmt.Printf("Переменная: %T\n , %v\n", b, b)

	fmt.Printf("Переменная: %T\n, %v\n", d, d)
	fmt.Printf("Переменная: %T\n, %v\n", r, r)

	// - var без типа
	var x2 = -12

	var f2 = 3.1415

	var name = "Serega"

	var false = false

	var d2 = 1234
	var r2 = 42

	fmt.Printf("Переменная: %T\n, %v\n", x2, x2)

	fmt.Printf("Переменная: %T\n, %v\n", f2, f2)

	fmt.Printf("Переменная: %T\n, %v\n", name, name)

	fmt.Printf("Переменная: %T\n, %v\n", false, false)

	fmt.Printf("Переменная: %T\n, %v\n", d2, d2)
	fmt.Printf("Переменная: %T\n, %v\n", r2, r2)

	// - :=
	x3 := 676

	f3 := float32(32.321)

	group := "PI-332"

	true := true

	d3 := uint(12)
	r3 := int32(123)

	fmt.Printf("Переменная: %T\n, %v\n", x3, x3)

	fmt.Printf("Переменная: %T\n, %v\n", f3, f3)

	fmt.Printf("Переменная: %T\n, %v\n", group, group)

	fmt.Printf("Переменная: %T\n, %v\n", true, true)

	fmt.Printf("Переменная: (%T)\n, %v\n", d3, d3)
	fmt.Printf("Переменная: %T\n, %v\n", r3, r3)

	task2()
}

// Задача 1.2

func task2() {
	// - через 3-ью переменную
	var x int = 10
	var y int = 20

	var z = x
	x = y
	y = z

	fmt.Printf("Переменная: %T\n , %v\n", x, x)
	fmt.Printf("Переменная: %T\n , %v\n", y, y)

	// - без 3-ей переменной
	var x2 int = 3
	var y2 int = 4
	x2, y2 = y2, x2
	fmt.Printf("Переменная: %T\n , %v\n", x2, x2)
	fmt.Printf("Переменная: %T\n , %v\n", y2, y2)
}
