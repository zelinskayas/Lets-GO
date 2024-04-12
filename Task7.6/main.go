package main

import "fmt"

func main() {
	sliceint := make([]int, 0, 10)
	sliceint = append(sliceint, 4, 1, 8, 9)

	fmt.Println(sliceint)
}
