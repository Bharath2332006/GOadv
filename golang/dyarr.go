package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	parts := strings.Fields(str)
	arr := []int{}
	for _, parts := range parts {
		nums, _ := strconv.Atoi(parts)
		arr = append(arr, nums)
	}
	fmt.Println(arr)
}
