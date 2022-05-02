package main

import(
	"fmt"
	"sort"
)

func main(){
	var numb=[]string{"a","b","c"}
	fmt.Println(numb)
	numb=append(numb,"d")
	fmt.Println(numb)
	numb=append(numb,"e")
	fmt.Println(numb)
	numb=append(numb[1:])
	fmt.Println(numb)
	numb=append(numb[1:len(numb)])
	fmt.Println(numb)
	numb=numb[1:len(numb)]
	fmt.Println(numb)
	numb=numb[:len(numb)-1]
	fmt.Println(numb)
	numbers:=make([]int,3,3)
	numbers[0]=3
	numbers[1]=2
	numbers[2]=5
	fmt.Println(numbers)
	numbers=append(numbers,1)
	fmt.Println(numbers,cap(numbers),len(numbers))
	sort.Ints(numbers)
	fmt.Println(numbers)
}