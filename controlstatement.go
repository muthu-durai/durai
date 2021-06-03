package main
import "fmt"
func main() {
for i:=0; i<5; i++{
fmt.Println(i)
if i == 3{
		break;
}
}
var x int = 0
Lable1: for x < 8 {
   if x == 5 {
       x = x + 1;
	  goto Lable1
   }
   fmt.Printf("value is: %d\n", x);
   x++;     
} 
}
