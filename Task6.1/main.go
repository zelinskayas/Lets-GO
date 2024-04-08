package main

import (
	"fmt"
)

type Contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	c1 := Contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v", c1)
}
