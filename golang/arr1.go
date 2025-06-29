package main
import (
	"fmt"
	"strings"
	"bufio"
	"strconv"
	"os"
)
func main(){
	arr:=[]int{}
	scanner :=bufio.NewScanner(os.Stdin) 
	scanner.Scan()
	line:= scanner.Text()
	parts:=strings.Fields(line)
	for _,parts:=range parts {
		nums,_:= strconv.Atoi(parts)
		arr=append(arr,nums)
	}
	fmt.Println(arr)
} 