package main
import "fmt"
func main() {
	var value string = "five"
switch value {
	case "one":
	fmt.Println("C#")
	case "two", "three":
	fmt.Println("Go")
	case "four", "five", "six":
	fmt.Println("Golang")
}
}
