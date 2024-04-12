package main

import "fmt"

func main() {
	sliceint1 := []int{1, 2, 3}
	sliceint2 := []int{4, 5, 6}

	sliceint := append(sliceint1, sliceint2...)

	fmt.Println(sliceint)
}
