package main

import "fmt"

func change(a *int) {
	*a = 66
}

func main() {
	b := 77
	fmt.Println("значение до вызова функции change: ", b)
	change(&b)
	fmt.Println("значение после вызова функции change: ", b)
}
