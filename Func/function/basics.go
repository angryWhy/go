package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(add3, "HAL"))
	//调用
	//显示的声明一个数组，将实参复制给这个数组
	values := []int{1, 2, 3}
	sum(values...)
	one(3.0, 3.2, 3.5)
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
func de() int {
	i := 9
	//这里是5，跟的是函数
	defer func() {
		fmt.Println(i)
	}()
	//这里是9，后面跟的是表达式，直接复制
	defer fmt.Println(i)
	return 5
}
func one(arr ...float64) (float64, error) {
	p := 1.0
	for _, v := range arr {
		p *= v
	}
	if p != 0 {
		return 1.0 / p, nil
	} else {
		return 0, errors.New("")
	}
}
