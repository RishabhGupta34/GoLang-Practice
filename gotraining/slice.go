package main
import "fmt"
func main(){
	a:=make([]int,3,3)
	fmt.Println(a)
	b:=append(a,3)
	b[0]=4
	fmt.Println(a,b)
}