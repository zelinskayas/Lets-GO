package main

import (
	"fmt"
	"log"
)

func main() {
	a := 12
	do(a)
}
func do(v any) {
	a := 10

	v1, ok := v.(int)
	if !ok {
		log.Fatalln("ошибка приведения к int")
	}
	a += v1

	fmt.Println(a)
}
