package main

import "fmt"

func main() {
	mapanimals := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	delete(mapanimals, "бегемот")
	fmt.Println(mapanimals)
}
