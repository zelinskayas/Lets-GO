package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type TempAmbient struct {
	mu  sync.RWMutex
	val string
}

func (t *TempAmbient) ReadTemp() string {
	defer t.mu.RUnlock()
	t.mu.RLock()
	return t.val
}

func (t *TempAmbient) ChangeTemp(v string) {
	defer t.mu.Unlock()
	t.mu.Lock()
	t.val = v
}

func main() {
	t := TempAmbient{}
	for i := 0; i < 12; i++ {
		j := i
		go func() {
			t.ChangeTemp(strconv.Itoa(j+7) + " градусов по цельсию")
		}()
		go func() {
			fmt.Printf("горутина: %v, температура окружающей среды: %v\n", j, t.ReadTemp())
		}()
	}
	time.Sleep(2 * time.Second)

	t.ChangeTemp("266 градусов по фаренгейту")
	fmt.Printf("main горутина, температура окружающей среды: %v", t.ReadTemp())
}
