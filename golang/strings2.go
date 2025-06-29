package main
import "fmt"
import "strings"
func main(){
	var s1 string
	fmt.Scan(&s1)
	fmt.Println(upper(s1))
	fmt.Println(lower(s1))
	fmt.Println(cont(s1))
	fmt.Println(replace(s1))
	fmt.Println(split(s1))
}
func upper(s string) string{
	return strings.ToUpper(s)
}
func lower(s string) string{
	return strings.ToLower(s)
}
func cont(s string) bool {
	if strings.Contains(s,"h"){
		return true
	}
	return false
}
func replace(s string ) string{
	return strings.ReplaceAll(s,"a","Z")
}
func split(s string) []string{
	return strings.Split(s,".")
}