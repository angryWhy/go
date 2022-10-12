package main

import (
	"fmt"
	"time"
)

func main() {
	test2()
}
func test() {
	//闭包
	arr := []int{1, 2, 3, 4, 5}
	for _, ele := range arr {
		go func() {
			fmt.Println(ele)
		}()
	}
	time.Sleep(1 * time.Second)
}
func test2() {
	//闭包
	arr := []int{1, 2, 3, 4, 5}
	for _, ele := range arr {
		go func(ele int) {
			fmt.Println(ele)
		}(ele)
	}
	time.Sleep(1 * time.Second)
}
