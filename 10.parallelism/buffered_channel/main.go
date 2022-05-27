package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	go func() { c <- 3 }()
	fmt.Println(<-c)
	time.Sleep(2 * time.Second)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
