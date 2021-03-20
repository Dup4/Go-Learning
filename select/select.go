package main

import (
	"fmt"
)

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Print("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Print("received ", i3, " from c3\n")
		} else {
			fmt.Print("c3 is closed\n")
		}
	default:
		fmt.Print("no communication\n")
	}
}
