package main

import (
	"fmt"
)

func main() {
	wordCount := make(map[string]int)
	var str string
	for {
		fmt.Scanln(&str)
		if str == "exit" {
			break
		}
		wordCount[str]++
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
