package main

import "fmt"

func main() {
	masstr := [4]string{"яблоко", "груша", "помидор", "абрикос"}

	//замена если заранее известен индекс
	//masstr[2] = "персик"

	//обход всех элементов с поиском помидора,
	//если нашли заменили на персик и вышли из цикла
	for i, _ := range masstr {
		if masstr[i] == "помидор" {
			masstr[i] = "персик"
			return
		}
	}

	fmt.Println(masstr)
}
