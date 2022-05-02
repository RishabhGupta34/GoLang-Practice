package main
import(
	"fmt"
)
func main(){
	var p *int
	if p!=nil{
	fmt.Println(*p)
	}else{
		fmt.Println("Null")
	}
	v:=43
	p=&v
	if p!=nil{
	fmt.Println(*p)
	}else{
		fmt.Println("Null")
	}
	*p=*p+4
	fmt.Println(*p)
	fmt.Println(v)
}