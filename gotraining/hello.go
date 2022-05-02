package main
import(
	"fmt"
)
func main(){
	var s=`
Hello,
			Hi`
	fmt.Printf("%s %t\n",s,false)
	fmt.Printf("%[2]d %[1]d %[2]d\n",10,20)
	fmt.Printf("'%*d'\n",10,200000) //10->size of padding
	i,j:=10,20
	i,j=j,i
	fmt.Println(i,j)
}