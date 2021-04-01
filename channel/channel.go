package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(60)) * time.Millisecond)
		channel <- "hello"
	}()

	msg := <-channel
	fmt.Printf(msg)
}
