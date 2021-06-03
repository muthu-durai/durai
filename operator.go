package main
import "fmt"
func main() {
	p := 23
	q := 60
	result1 := p + q
	fmt.Printf("Result of p + q = %d\n", result1)
	result2 := p == q
	fmt.Println(result2)
	result3 := p != q
	fmt.Println(result3)
	if p != q && p <= q {
		fmt.Println("True")
	}

	if p != q || p <= q {
		fmt.Println("True")
	}

	if !(p == q) {
		fmt.Println("True")
	}
	result4 := p & q
	fmt.Printf("Result of p & q = %d\n", result4)
	p = q
	fmt.Println(p)

}
