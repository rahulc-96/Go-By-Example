package main  
import "fmt"
func main(){
	 m:=map[string]string{"Ana":"A+","James":"B"}
	 for k,v:=range m{
	 	fmt.Println(k,"->",v)
	 }
     i:=0
	 for k:=range m{
        i++
	 	fmt.Println("Key",i,k)
	 }
     i=0
	 for _,v:=range m{
        i++
	 	fmt.Println("Value",i,v)
	 }

	 arr:=[]int{1,2,3,4}
	 for i,v:=range arr{
	 fmt.Println("Element at",i,"is",v)}
}