package main

import "fmt"

func main() {
	sliceint1 := []int{1, 2, 3, 4, 5, 6}

	//индекс искомого числа и искомое число, что нужно удалить
	var ind, valx int
	valx = 4

	for i, val := range sliceint1 {
		if val == valx {
			ind = i
		}
	}

	sliceint := sliceint1[0:ind]
	sliceint = append(sliceint, sliceint1[ind+1:]...)

	fmt.Println(sliceint)
}
