package main

import (
	"fmt"
	"os"
)

func main() {
	data := `class main{
		public static void main(String[] args){
		System.out.println("Hello World");
		}
		}`
	err := os.WriteFile("w1.java", []byte(data), 0644)
	if err != nil {
		fmt.Println("error in write", err)

	} else {
		fmt.Println("file written successfully")
	}
}
