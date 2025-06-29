package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(sum(a, b))
}
func sum(a int, b int) int {
	return a + b
}
