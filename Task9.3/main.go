package main

import "fmt"

func checkFood(n string) {
	switch n {
	case "груша", "яблоко", "апельсин":
		fmt.Println("это фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println("это овощ")
	default:
		fmt.Println("что-то странное...")
	}
}

func main() {
	checkFood("огурец")
	checkFood("яблоко")
	checkFood("вишня")
}
