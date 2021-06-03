package main
import "fmt"
func main() {
	var arr = [7]string{"This", "is", "the", "program","of", "Go", "lang"}
	fmt.Println("Array:", arr)
	myslice := arr[1:6]
	fmt.Println("Slice:", myslice)
	fmt.Printf("Length of the slice: %d", len(myslice))
	fmt.Printf("\nCapacity of the slice: %d\n", cap(myslice))
}