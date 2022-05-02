package main

import(
	"fmt"
)

func main(){
	var abc[3] string
	abc[1]="hello"
	abc[2]="hi"
	fmt.Println(abc)

	var numb=[5]int{5,3,2,3,5}
	fmt.Println(numb)
	fmt.Println(len(abc))
	fmt.Println(len(numb))
	fmt.Println(numb)
}