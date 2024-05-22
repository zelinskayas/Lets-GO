package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	d := time.Now().Add(time.Second * 2)
	ctx, cancel := context.WithDeadline(ctx, d)
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("стоп горутина:", i, ctx.Err())
					return
				default:
					fmt.Println("горутина: ", i, " чем-то занята")
					time.Sleep(time.Second)
				}
			}
		}()
	}
	wg.Wait()
}
