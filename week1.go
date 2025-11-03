package main
import "fmt"
func gcd(a int,b int) int{
	for b!=0{
		a , b = b , a % b
	}
	return a
}
func lcm(a int,b int) int{
	return(a*b)/gcd(a,b)
}
func main(){
	var a,b int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a,&b)
	gcdResult:=gcd(a,b)
	lcmResult:=lcm(a,b)
	fmt.Printf("The gcd of %d and %d is %d\n", a,b,gcdResult)
	fmt.Printf("The lcm of %d and %d is %d\n", a,b,lcmResult)
}