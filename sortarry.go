package main
import "fmt"
func main(){
	var n int
	fmt.Scanln(&n)
	arr:=make([]int,n)
	for i:=0;i<n;i++ {
		fmt.Scanln(&arr[i])
	}
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if arr[i]>arr[j]{
				temp:=arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
		}
	}
	fmt.Println(arr)
}