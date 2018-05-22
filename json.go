package main
import "fmt"
import "encoding/json"
type book struct{
        Name string
        Author string
	}

func main(){

	

	fmt.Println(book{Name:"Game Of Thrones",Author:"GRRM"})
	sp:=book{Name:"Harry Potter",Author:"JKR"}
	ptr:=&sp
	fmt.Println("Author of "+ptr.Name+" is "+ptr.Author)
	s:=book{Name:"Game of Thrones",Author:"GRRM"}
    sj:=&book{Name:"Naruto",Author:"MAS"}
    js,_:=json.Marshal(sj)
	fmt.Println(string(js))
	fmt.Println("Title: "+s.Name+" Author: "+s.Author)

}

