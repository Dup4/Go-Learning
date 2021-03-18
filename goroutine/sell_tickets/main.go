package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var total_tickets int32 = 20
var mutex = &sync.Mutex{}

func sell_tickets(i int, done chan bool) {
	for {
		if total_tickets > 0 { // 如果有票就卖
			mutex.Lock()
			if total_tickets > 0 {
				time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
				total_tickets-- // 卖一张票
				fmt.Println("id:", i, "  ticket:", total_tickets)
			}
			mutex.Unlock()
		} else {
			done <- true
			break
		}
	}
}

func main() {
	done := make(chan bool)
	n := 8
	runtime.GOMAXPROCS(n)        // 我的电脑是 8 核处理器，所以我设置了 8
	rand.Seed(time.Now().Unix()) // 生成随机种子
	for i := 0; i < n; i++ {     // 并发 8 个 goroutine 来卖票
		go sell_tickets(i, done)
	}

	for i := 0; i < n; i++ {
		<-done
	}
	fmt.Println(total_tickets, "done") // 退出时打印还有多少票
}
