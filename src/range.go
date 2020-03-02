package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // 關閉 Channel
	}()
	for i := range c { // 在 close 後跳出迴圈
		fmt.Println(i)
	}
}
