package main
import "fmt"
func main(){
	a:=make([]int,2000,2000)
	a=append(a,9)
	fmt.Println(cap(a))
}