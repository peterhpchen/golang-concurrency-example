package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0
	ch := make(chan int, 1)
	ch <- total
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- <-ch + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
