package main
import "fmt"
type Person struct{
	age int
}
func (p *Person) SetAge(age int) {
	p.age=age
}
func (p Person) IsMajor() bool{
	return p.age>=18
}

func main(){
	p:=Person{}
	p.SetAge(117)
	fmt.Println(p.IsMajor())
}