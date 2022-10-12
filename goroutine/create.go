package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func add1() {
	defer wg.Done()
	fmt.Println("over")
	go sub2()
}
func sub2() {
	fmt.Println("over2")
}
func main() { //main协程
	wg.Add(2)
	//开启协程
	go add1()
	go add1()
	wg.Wait()
}
