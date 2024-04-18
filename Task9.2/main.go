package main

import "fmt"

func main() {
	sliceint := []int{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Println("v1:", sliceint[i])
	LABEL:
		for j := 0; j < 3; j++ {
			fmt.Println("\tv2:", sliceint[j])
			for n := 0; n < 3; n++ {
				fmt.Println("\t\tv3:", sliceint[n])
				for k := 0; k < 3; k++ {
					fmt.Println("\t\t\tv4:", sliceint[k])
					if k == 1 {
						break LABEL
					}
				}
			}
		}
	}
}
