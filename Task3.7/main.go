package main

import "fmt"

func main() {
	const (
		const1 = 1 << iota
		const2
		const3
		const4
		const5
	)

	fmt.Println(const1, const2, const3, const4, const5)

}
