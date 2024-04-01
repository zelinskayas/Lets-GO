package main

import "fmt"

func hello(f func()) {
	f()
}

func main() {
	anonf := func() {
		fmt.Println("Hello, Go!")
	}

	hello(anonf)
}
