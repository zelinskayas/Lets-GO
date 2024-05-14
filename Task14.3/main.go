package main

import "fmt"

func main() {
	ch := make(chan string, 4)
	ch <- "Привет"
	ch <- "буферизованный канал!"

	for i := 0; i <= len(ch); i++ {
		fmt.Println(<-ch)
	}
}
