package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(add3, "HAL"))
	//调用
	//显示的声明一个数组，将实参复制给这个数组
	values := []int{1, 2, 3}
	sum(values...)
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
func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
