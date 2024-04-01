package main

import "fmt"

func hello() {
	fmt.Println("Hello, Go!")
}

func main() {
	defer fmt.Println("завершение работы")
	hello()
}
