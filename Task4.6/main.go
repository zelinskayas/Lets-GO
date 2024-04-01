package main

import "fmt"

func fib(a, b, c int) {
	if c == 0 {
		return
	}
	fmt.Println(a)
	fib(b, a+b, c-1)
}

func main() {
	fib(0, 1, 23)
}
