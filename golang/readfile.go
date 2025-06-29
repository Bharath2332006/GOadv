package main
import (
	"fmt"
	"os"
)
func main(){
	content,err :=os.ReadFile("w1.java")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("readed")
		fmt.Println(string(content))
	}

}