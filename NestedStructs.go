package main
import "fmt"
 type details struct{
        	isbn string
        	rel  string
        }

type book struct{
        Name string
        Author string
        details
	}

func main(){

	
	s:=book{Name:"Game Of Thrones",Author:"GRRM",details:details{isbn:"IND5113",rel:"25/12/1998"}}
	fmt.Println(s)
	fmt.Println("ISBN of "+s.Name+" is "+s.details.isbn)

}

