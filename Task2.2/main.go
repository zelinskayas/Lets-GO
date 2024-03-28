package main

import "fmt"

func main() {
	var a, b, c, d int
	a = 16
	b = 3
	c = a / b
	d = a % b

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T.", c, d, c)
}
