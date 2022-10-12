package main

import "fmt"

func main() {

	go test5()
}
func test4() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	panic("oops")
	defer fmt.Println("9")
	fmt.Println("none")
}
func test5() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer func() {
		//中途结束，不影响其他协程
		recover()
	}()
	defer fmt.Println("3")
	panic("oops")
}
