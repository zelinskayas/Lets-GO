package main

import "fmt"

type square int

func StringSquare(s square) {
	fmt.Println(s, "м²")
}

func main() {
	var s square = 34
	s = s + 10
	StringSquare(s)
}
