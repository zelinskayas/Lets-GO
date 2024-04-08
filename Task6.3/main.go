package main

import (
	"fmt"
)

type Contract struct {
	ID     int
	Number string
	Date   string
}

func (c *Contract) String() string {
	return "Договор № " + c.Number + " от " + c.Date
}

func main() {
	c1 := Contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Println(c1.String())
}
