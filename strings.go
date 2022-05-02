package main
import(
	"fmt"
	"strings"
)
func main(){
	str:="This is go lang"
	fmt.Println(str)
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToTitle(str))
	fmt.Println(strings.Title(str))
	fmt.Println(strings.ToLower(str))
	str1:="hello"
	str2:="Hello"
	fmt.Println("Case sensitive:",str1==str2)
	fmt.Println("Case non sensitive:",strings.EqualFold(str1,str2))
	fmt.Println("Contains hell:",strings.Contains(str,"hell"))
	fmt.Println("Contains hell:",strings.Contains(str1,"hell"))
	fmt.Println("Contains hell:",strings.Contains(str2,"hell"))
}