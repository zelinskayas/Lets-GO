package main

import "fmt"

const testConst = "глобальная константа"

func main() {
	const testConst = "затенили локальной константой"
	fmt.Println(testConst)
}
