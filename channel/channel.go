package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string, 2)

	go func() {
		time.Sleep(time.Duration(rand.Intn(60)) * time.Millisecond)
		channel <- "hello"
		channel <- "hello2"
	}()

	msg := <-channel
	msg2 := <-channel
	fmt.Printf(msg, msg2)
}
