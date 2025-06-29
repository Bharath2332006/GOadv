package main

import (
	"fmt"
	"time"
)

func addandsub(a int, b int) {
	fmt.Println(a+b, a-b)
	time.Sleep(10 * time.Second)
}
func main() {
	//cls
	// var a, b int
	for i := 0; i < 3; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		go addandsub(a, b)
		addandsub(a, b)
		time.Sleep(2 * time.Second)
	}
}
