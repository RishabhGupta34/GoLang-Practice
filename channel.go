package main
import(
	"fmt"
)
var n=[]string{"new1","new2","new3"}
var v=[]string{"vid1","img","vid3"}
var i=[]string{"img","img","img"}
var a=[]string{"all1","all2","all3"}

func News(c chan string,find string){
	for _,ele := range n{
		if find==ele{
			c<-find
		}
	}
	close(c)
}
func Vid(c chan string,find string){
	for _,ele := range v{
		if find==ele{
			c<-find
		}
	}
	close(c)
}
func Imag(c chan string,find string){
	for _,ele := range i{
		if find==ele{
			c<-find
		}
	}
	close(c)
}
func All(c chan string,find string){
	for _,ele := range a{
		if find==ele{
			c<-find
		}
	}
	close(c)
}
func main(){
	news:=make(chan string,10)
	vid:=make(chan string,10)
	imag:=make(chan string,10)
	all:=make(chan string,10)
	go News(news,"img")
	go Vid(vid,"img")
	go Imag(imag,"img")
	go All(all,"img")

	for{
		select{
		case ne,more:=<-news:
			if !more{
				news=nil
			}else{
				fmt.Println("News",ne)
			}
		case vi,more:=<-vid:
			if !more{
				vid=nil
			}else{
				fmt.Println("Video",vi)
			}
		case im,more:=<-imag:
			if !more{
				imag=nil
			}else{
				fmt.Println("Image",im)
			}
		case al,more:=<-all:
			if !more{
				all=nil
			}else{
				fmt.Println("All",al)
			}
		}
		if news==nil && vid==nil && imag==nil && all==nil{
			break
		}
	}


}