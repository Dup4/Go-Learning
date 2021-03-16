package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func f(msg string) {
	fmt.Println(msg)
}

func routine(name string, delay time.Duration) {
	t0 := time.Now()
	fmt.Println(name, " start at ", t0)
	time.Sleep(delay)
	t1 := time.Now()
	fmt.Println(name, " end at ", t1)
	fmt.Println(name, " lasted ", t1.Sub(t0))
}

func input_wait() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	runtime.GOMAXPROCS(5)
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	rand.Seed(time.Now().Unix())

	var name string
	for i := 0; i < 3; i++ {
		name = fmt.Sprintf("go_%02d", i)
		go routine(name, time.Duration(rand.Intn(5))*time.Second)
	}

	input_wait()
	fmt.Println("done")
}
