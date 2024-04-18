package main

import "fmt"

func fruitMarket(f string) {
	fruits := make(map[string]int, 10)
	fruits["апельсины"] = 5
	fruits["яблоки"] = 3
	fruits["сливы"] = 1
	fruits["груши"] = 0

	val, ok := fruits[f]
	if ok {
		fmt.Printf("\"%s\" - %d", f, val)
	} else {
		fmt.Println("такой фрукт отсутствует")
	}

}

func main() {
	fruitMarket("апельсины")
}
