package main

import "fmt"

func main() {
	mapanimals := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	mapanimals["бегемот"] = 2

	fmt.Println(mapanimals)
}
