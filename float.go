package main
import(
	"fmt"
	"math/big"
)
func main(){
	f1,f2,f3 := 2.6,3.6,4.2
	i1:=3
	fmt.Println(f1+f2+f3)
	fmt.Println(float64(i1)+f1+f2+f3)
	var bf1,bf2,bf3,bs big.Float
	bf1.SetFloat64(2.6)
	bf2.SetFloat64(2.6)
	bf3.SetFloat64(2.6)
	bs.Add(&bf1,&bf2).Add(&bs,&bf3) // & required for bigfloat addition
	fmt.Printf("%.10g\n",&bs) // & required for bigfloat print
}