package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("calculate goroutine starts calculating")
		time.Sleep(time.Second) // Heavy calculation
		fmt.Println("calculate goroutine ends calculating")

		ch <- "FINISH"

		fmt.Println("calculate goroutine finished")
	}()

	fmt.Println("main goroutine is waiting for channel to receive value")
	fmt.Println(<-ch) // goroutine 執行會在此被迫等待
	fmt.Println("main goroutine finished")
}
