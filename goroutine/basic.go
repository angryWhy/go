package main

import (
	"fmt"
	"sync"
)

func main() {
	go add(2, 3)
	go func(a, b int) {
		fmt.Println(a + b)
	}(2, 2)
}
func add(a, b int) {
	v := a + b
	fmt.Println(v)
}
func gor() {
	//sync包
	wg := sync.WaitGroup{}
	//计数器：10个协程
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(a, b int) {
			defer wg.Done()
			//do something
		}(i, i+1)
	}
	//等待为0
	wg.Wait()
}
