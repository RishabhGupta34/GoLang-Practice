package main
import(
	"fmt"
)

type Vehicle struct{
	name string
	Color string
}

func main(){
	a:=Vehicle{"Alto","red"}
	fmt.Println(a)
	fmt.Printf("%+v\n",a)
	fmt.Println(a.Color)
}