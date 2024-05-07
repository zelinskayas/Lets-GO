package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int
	Landlord string
	tenat    string
}

func main() {
	c := contract{
		Number:   2,
		Landlord: "Остап",
		tenat:    "Шура",
	}

	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", string(res))
}
