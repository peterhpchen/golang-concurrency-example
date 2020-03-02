package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	v   int
	mux sync.Mutex // 互斥鎖
}

func main() {
	total := SafeNumber{v: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total.mux.Lock()
			total.v++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total.mux.Lock()
	fmt.Println(total.v)
	total.mux.Unlock()
}
