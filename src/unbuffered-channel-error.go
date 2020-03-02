package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 1 // 等到天荒地老
	fmt.Println(<-ch)
}
