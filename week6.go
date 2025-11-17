package main
import "fmt"
func main () {
	fmt.Println("Enter first String:")
	var first string
	fmt.Scanln(& first)
	fmt.Print("Enter second String:")
	var second string
	fmt.Scanln(& second)
	fmt.Print("Concatenation of two strings : ", first + second)
}