package main  
import "fmt"
func main(){
  var a[5] int
  fmt.Println("Default Array a:",a)

  b:=[5]int{1,2,3,4,5}
  fmt.Println("Intialized Array b:",b)

  a[2]=10001
  fmt.Println("Modified Array a:",a)

}