package main

import "fmt"

func main() {
	a, b, c, d := 4, 9, 8, 34

	fmt.Println("значения: ", a, b, c, d, " адреса: ", &a, &b, &c, &d)
}
