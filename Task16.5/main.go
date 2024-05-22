package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	chstop := make(chan struct{})
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-chstop:
					fmt.Println("stop горутина: ", i)
					return
				default:
					time.Sleep(time.Second)
					fmt.Println("сложные вычисления горутины: ", i)
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		fmt.Println("ой, всё!")
		close(chstop)
	}()

	wg.Wait()
}
