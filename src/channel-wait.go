package main

import (
	"fmt"
	"time"
)

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	c <- "FINISH"
}

func main() {
	ch := make(chan string)

	go say("world", ch)
	go say("hello", ch)

	<-ch
	<-ch
}
