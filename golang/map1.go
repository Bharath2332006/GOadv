package main
import "fmt"

func main(){
	day:= map[int]string{
		1:"MON",
		2:"TUE",
		3:"WED",
		4:"THU",
		5:"FRI",
		6:"SAT",
		7:"SUN",
	}
	fmt.Println(day)
	fmt.Println(day[7])
}