package main

import (
	"fmt"
	"strings" )

	func main(){
		var str string
		fmt.Scanln(&str)
		words := strings.Fields(str)
		wordCount := make(map[string]int)
		for _, word := range words {
			word = strings.ToLower(word)
			wordCount[word]++
		}

		for word, count := range wordCount {
			fmt.Printf("%s: %d\n", word, count)
		}
	}

