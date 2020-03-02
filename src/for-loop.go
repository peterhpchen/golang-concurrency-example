package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // 關閉 Channel
	}()
	for {
		v, ok := <-c
		if !ok { // 判斷 Channel 是否關閉
			break
		}
		fmt.Println(v)
	}
}
