package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() { // calculate goroutine
		fmt.Println("calculate goroutine starts calculating")
		time.Sleep(time.Second) // Heavy calculation
		fmt.Println("calculate goroutine ends calculating")

		ch <- "FINISH" // goroutine 執行會在此被迫等待

		fmt.Println("calculate goroutine finished")
	}()

	time.Sleep(2 * time.Second) // 使 main 比 goroutine 慢
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println("main goroutine finished")
}
