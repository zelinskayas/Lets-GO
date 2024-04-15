package main

import "fmt"

func main() {
	mapanimals := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	fmt.Println(mapanimals)
}
