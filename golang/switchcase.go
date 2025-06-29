package main

import "fmt"

func main() {
	var day int
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("MONDAY")
	case 2:
		fmt.Println("TUESDAY")
	case 3:
		fmt.Println("WEDNESDAY")
	case 4:
		fmt.Println("THURSDAY")
	case 5:
		fmt.Println("FRIDAY")
	case 6:

		fmt.Println("SATURDAY")
	case 7:
		fmt.Println("SUNDAY")
	}
}
