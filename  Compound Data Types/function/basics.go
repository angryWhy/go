package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(add3, "HAL"))
}

//	func name(params)(result-list){
//		body
//	}
//
// 返回值
func add(x, y int) int {
	return x + y
}
func sub(x, y int) (z int) {
	z = x + y
	return
}
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func add3(item rune) rune {
	return item + 1
}
