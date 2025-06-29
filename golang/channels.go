package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go task(ch)
	fmt.Println("waiting for task")
	for val:= range ch{
		
		fmt.Println(val)
	}
}
func task(ch chan int) {
	fmt.Println("task going")
	time.Sleep(3 * time.Second)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}
