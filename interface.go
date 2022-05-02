package main
import "fmt"
type Speaker interface{
	giveLang() string
}
type Cat struct{
	lang string
}
type Dog struct{
	lang string
}
type Human struct{
	lang string
}
type Alien struct{
	lang string
}
func (c Cat) giveLang() string{
	return c.lang
}
func (c Cat) giveL() string{
	return c.lang
}
func (d Dog) giveLang() string{
	return d.lang
}
func (h Human) giveLang() string{
	return h.lang
}
func (a Alien) giveLang() string{
	return a.lang
}
func speak(s Speaker){
	fmt.Println(s,s.giveLang())
}
func main(){
	c:=Cat{"Meow"}
	d:=Dog{"Bow"}
	h:=Human{"Hi"}
	a:=Alien{"wedsvc"}
	speak(c)
	speak(d)
	speak(h)
	speak(a)
	fmt.Println("\n\n")

	all:=make([]Speaker,0,5)
	all=append(all,Cat{"Meow"})
	all=append(all,Dog{"Bow"})
	fmt.Println(all[0].giveLang())
	// fmt.Println(all[0].giveL())
	for _,v:=range all{
		speak(v)
	}
	fmt.Printf("%#v\n",all)
}