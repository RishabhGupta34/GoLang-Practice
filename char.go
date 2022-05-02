package main

import(
	"fmt"
)
func main(){
	str1:="Hi"
	str2:="i"
	str3:="am"
	strlength,err:=fmt.Println(str1,str2,str3)
	if(err==nil){
		fmt.Println("Length:",strlength)
	}
	num:=43
	boo:=true
	fmt.Printf("%v\n",num)
	fmt.Printf("%v\n",boo)
	fmt.Printf("%v\n",str1)
	fmt.Printf("%.2f\n",float64(num))
	fmt.Printf("%T,%T,%T,%T\n",num,boo,str1,float64(num))
	str4:=fmt.Sprintf("%T,%T,%T,%T\n",num,boo,str1,float64(num))
	fmt.Println(str4)

}
