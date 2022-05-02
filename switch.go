package main
import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())
	// d:=rand.Intn(5)+1
	switch d:=rand.Intn(5)+1;d{
		case 1:
			fmt.Println("O")
		case 3:
			fmt.Println("T")
		default:
			fmt.Println("D")
	}
	x:=-43
	switch{
		case x<0:
			fmt.Println("<O")
		case x>0:
			fmt.Println(">0")
		default:
			fmt.Println("==0")
	}
}