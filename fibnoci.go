package main
import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
	for i:=0;i<n;i++{
	fmt.Print(fib(i))
	}
}
func fib(n int) int{
	if n<=1 {
		return n
	}
	return fib(n-1)+fib(n-2)
}