package main

import "fmt"

func main() {
	mapanimals := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	searchanim := []string{"слон", "бегемот", "осьминог"}

	for _, val := range searchanim {
		valcnt, ok := mapanimals[val]
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", val, valcnt, ok)
	}
}
