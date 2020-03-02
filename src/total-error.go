package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0
	for i := 0; i < 1000; i++ {
		go func() {
			total++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(total)
}
