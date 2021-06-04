package main
import("fmt"
"time")
	func portal1(channel1 chan string) {

		time.Sleep(3*time.Second)
		channel1 <- "Welcome to channel 1"
	}
	func portal2(channel2 chan string) {

		time.Sleep(9*time.Second)
		channel2 <- "Welcome to channel 2"
	}
func main(){
R1:= make(chan string)
R2:= make(chan string)
go portal1(R1)
go portal2(R2)
select{
	case op1:= <- R1:
	fmt.Println(op1)
	case op2:= <- R2:
	fmt.Println(op2)
}
}
