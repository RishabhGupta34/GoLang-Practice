package main
import(
	"fmt"
)
func main(){
	hello()
	add(3223,40945)
	sum:=add1(3223,40945)
	fmt.Println(sum)
	add2("Hey",3223,40945)
	sum=addAll(3223,40945,234,453)
	fmt.Println(sum)
	full,sum:=returnMult("rishabh","gupta")
	fmt.Println(full,sum)
	full,sum=returnMult("rishabh","gupta")
	fmt.Println(full,sum)
}
func hello(){
	fmt.Println("Hello")
}
func add(val1 int,val2 int){
	fmt.Println(val1,val2,val1+val2)
}
func add1(val1,val2 int) int{
	return val1+val2
}
func add2(s string,val1,val2 int){
	fmt.Println(s,val1,val2)
}
func addAll(values ... int) int{
	sum:=0
	for i:=range values{
		sum+=values[i]
	}
	return sum
}
func returnMult(f,l string) (string,int){
	full:=f+" "+l
	return full,len(full)
}
func returnMultNew(f,l string) (full string,length int){
	full=f+" "+l //No need to declare because declared in return type
	length=len(full)
	return 
}