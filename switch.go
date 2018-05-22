package main 
import "fmt"
import "time"
func main(){
	t:=time.Now()
	switch{
		case t.Hour()<12:fmt.Println("Good Morning")
		               break
	    case t.Hour()>12&&t.Hour()<4:fmt.Println("Good Afternoon")
	                             break
	    case t.Hour()>5&&t.Hour()<7:fmt.Println("Good Evening")
	                             break
	    default:fmt.Println("Good Night")                         
	}
}