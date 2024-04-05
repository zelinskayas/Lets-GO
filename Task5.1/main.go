package main

import "fmt"

func main() {
	s := "строковое значение"
	var pointerStr *string
	pointerStr = &s

	fmt.Println("указатель на строковое значение: ", pointerStr)
}
