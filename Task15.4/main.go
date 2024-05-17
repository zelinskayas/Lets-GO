package main

import (
	"fmt"
	"sync"
)

var loadOnce sync.Once

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			loadOnce.Do(start)
			fmt.Println("горутина: ", i)
		}()
	}
	wg.Wait()
}
func start() {
	fmt.Println("hi, func start")
}
