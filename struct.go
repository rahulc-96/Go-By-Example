package main
import "fmt"
type book struct{
        name string
        author string
	}


func main(){

	

	fmt.Println(book{name:"Game Of Thrones",author:"GRRM"})
	sp:=book{name:"Harry Potter",author:"JKR"}
	ptr:=&sp
	fmt.Println("Author of "+ptr.name+" is "+ptr.author)
	s:=book{name:"Game of Thrones",author:"GRRM"}
	fmt.Println("Title: "+s.name+" Author: "+s.author)
}

