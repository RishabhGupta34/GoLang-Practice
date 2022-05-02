package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	var s string ="helloo"
	const fh=43
	const fg int=43
	fmt.Println(s,s)
	fmt.Scanln(&s)
	fmt.Print(s,s)
	fmt.Print("\n")
	fmt.Print('\n')
	fmt.Print("\n")
	fmt.Println(s,s)
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	fmt.Println(str)
	fmt.Print("Enter num: ")
	str,_ = reader.ReadString('\n')
	f,err:=strconv.ParseFloat(strings.TrimSpace(str),64)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("value:",f)
	}
	fmt.Println("Hello")
}