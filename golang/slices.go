package main

import "fmt"

func main() {
	var arr []int
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	fmt.Println(arr)
	fmt.Println(arr[1:5])
}
 