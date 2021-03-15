package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 10
	var str string = "hello world"
	var i, j, k int = 1, 2, 3

	var b = true

	y := 100

	const _str string = "hello world"
	const PI float32 = math.Pi

	fmt.Println(x, str, i, j, k, b, y, _str, PI)
}
