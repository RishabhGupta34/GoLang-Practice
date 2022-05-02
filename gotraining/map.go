package main
import "fmt"
func main(){
	m:=map[string]string{
		"India":"Delhi",
	}
	for k,v:=range m{
		fmt.Println(k,v)
	}
	v,ok:=m["USA"]
	fmt.Println(v,ok)

	m2:=map[string] map[string]string{
		"India":map[string]string{
			"Haryana":"Chandigarh",
		},
	}
	m2["USA"]=map[string]string{
		"ABV":"CDF","DCDS":"sdc"}
	for k1,v1:=range m2{
		for k2,v2:=range v1{
			fmt.Println(k1,k2,v2)
		}
	}
	v2,ok:=m2["USA"]
	fmt.Println(v2,ok)
}