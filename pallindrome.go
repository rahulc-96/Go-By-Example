package main
import "fmt"
func main() {

	var num int
	fmt.Println("Enter a postive number")
	fmt.Scanf("%d",&num)
	var d=num
	var rev int
	for num>0{
		var digit=num%10
		rev=10*rev+digit
		num=num/10		
	}
	if rev==d{
		fmt.Println("Yes Number is Pallindrome")

	}else{
		fmt.Println("No Number is not Pallindrome")
	}
	
}