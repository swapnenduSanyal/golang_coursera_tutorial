/*
Here the race condition occurs because the order in which the values are
printed is not deterministic. It depends on which program finishes earlier
*/

package main

import	"fmt"
import "time"

func Double(x int){
	for i:= 0; i<100;i++{}
	fmt.Println("Double ",2*x)
}
func Increment(x int){
	for i:= 0; i<100;i++{}
	fmt.Println("Increment ",x+1)
}

func main() {
	x := 5
	go Double(x)
	go Increment(x)
	fmt.Println("Main ",x)
	<-time.After(time.Second*1)
}
