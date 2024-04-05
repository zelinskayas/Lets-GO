package main

import "fmt"

func main() {
	//переменная
	a := 5

	//указатель на переменную
	pointerInt := &a

	fmt.Println("До: ", a, pointerInt)

	//поменяли значение переменной через указатель
	*pointerInt = 7

	fmt.Println("После: ", a, pointerInt)
}
