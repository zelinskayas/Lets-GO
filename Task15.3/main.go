package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	var mu sync.Mutex
	for i := 0; i < 777; i++ {
		go func() {
			defer mu.Unlock()
			mu.Lock()
			n++
		}()
	}
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(n)
}
