package main

import "fmt"

func main() {
	chstr := make(chan string)
	go func() {
		chstr <- "Привет, строковый канал!"
	}()
	fmt.Println(<-chstr)
}
