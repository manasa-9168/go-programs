package main
import "fmt"
func main(){
	var rows int
	fmt.Print("Enter a number of rows:")
	fmt.Scan(&rows)
	for i := 1;i <= rows;i++{
		for j := 1;j <= rows-i;j++{
			fmt.Print(" ")
		}
		for k := 1;k <= i;k++{
			fmt.Print(k)
		}
		for l := i-1;l >= 1;l--{
			fmt.Print(l)
		}
		fmt.Println()
	}
}