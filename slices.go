package main 
import "fmt"
func main(){
	s:=make([]string,3)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("Intial Slice:",s)

	s=append(s,"d")
	s=append(s,"e","f")

	fmt.Println("After appending:",s)

	fmt.Println("Slicing1:",s[0:3])
	fmt.Println("Slicing2:",s[2:])
	fmt.Println("Slicing3:",s[:5])
}