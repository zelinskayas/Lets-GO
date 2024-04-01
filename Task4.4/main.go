package main

import "fmt"

func hello() func() {
	return func() {
		fmt.Println("Hello, Go!")
	}
}

func main() {
	f := hello()
	f()
}
