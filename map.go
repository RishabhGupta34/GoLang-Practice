package main

import(
	"fmt"
	"sort"
)

func main(){
	m:=make(map[string]int)
	m["bc"]=2
	m["ade"]=5
	m["efg"]=1
	fmt.Println(m)
	fmt.Println(m["bc"])
	fmt.Println(m["abcv"])
	delete(m,"bc")
	fmt.Println(m["bc"])
	m["bc"]=m["bc"]+3
	fmt.Println(m["bc"],"\n\n")

	for key,val := range m{
		fmt.Printf("%v: %v\n",key,val)
	}
	keys:=make([]string,len(m))
	i:=0
	for key:=range m{
		keys[i]=key
		i++
	}
	fmt.Println(keys,"\n\n")
	sort.Strings(keys)
	for j:=range keys{
		fmt.Println(keys[j])
	}

	values:=make([]int,len(m))
	i=0
	for _,value:=range m{
		values[i]=value
		i++
	}
	fmt.Println(values)

}