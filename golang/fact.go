package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println(fact(a))
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
