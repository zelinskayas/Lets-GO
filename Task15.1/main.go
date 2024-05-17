package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i < 6; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("горутина: ", j)
		}()
	}
	wg.Wait()
}
