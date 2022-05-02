package main 
import(
	"fmt"
	"errors"
)
func oper(x,y int) (int,int,int,int,error){
	if y==0{
		return 0,0,0,0,errors.New("Divide by zero")
	}
	return x+y,x-y,x*y,x/y,nil
}

func main(){
	a,s,m,d,err:=oper(5,9)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(a,s,m,d)
	}
}