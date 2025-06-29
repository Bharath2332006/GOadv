package main
import(
	"fmt"
	"os"
	"strings"
	"sync"
	"log"
	"bufio"
)
func main(){
	file:=[]string{"server1.log","server2.log","server3.log"}
	err:= processlog(file,"error1.log")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("success")
}
func processlog(input []string,output string)error{
	errs:=make(chan string)
	var wg sync.WaitGroup
	for _,f:= range input {
		wg.Add(1)
		go func(file string){
			defer wg.Done()
			readerror(file,errs)
		}(f)
	}
	go func(){
		wg.Wait()
		close(errs)
	}()
	return writeerror(output,errs)
	
}
func readerror(file string,errs chan <-string){
	f,err:=os.Open(file)
	if err!=nil{
		log.Println("open err:",err)
		return
	}
	defer f.Close()
	sc:=bufio.NewScanner(f)
	for sc.Scan(){
		if strings.Contains(sc.Text(),"ERROR"){
			errs<-sc.Text()
		}
	}
}
func writeerror(file string,errs <-chan string)error{
	f,err:=os.Create(file)
	if err!=nil{
		return err
	}
	defer f.Close()
	w:=bufio.NewWriter(f)
	for line:=range errs{
		_,err:=w.WriteString(line+"\n")
		if err!=nil{
			return err
		}
	}
	return w.Flush()
}