package main
import (
	"fmt"
	"os"
	//"strings"
)
func main() {
  f, err:=os.OpenFile(("W1.txt"),os.O_APPEND|os.O_WRONLY,0644)
   if err!=nil{
	fmt.Println(err)
	return
   }
   defer f.Close()
   content, err := f.WriteString("\nNEE THA BHARATHAAA")
   if err != nil {
	fmt.Println(err)
   }else{
	fmt.Println(content)
   }
}
