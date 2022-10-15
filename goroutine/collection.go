package main

import (
	"fmt"
	"sync"
)

type user struct {
	name string
}

var (
	lst []int = make([]int, 5)
	arr [5]int
	mp  sync.Map
	wg  = sync.WaitGroup{}
)

func main() {
	wg.Add(3)
	go rs()
	go rs2()
	go ms()
	wg.Wait()
	fmt.Println(mp.Load(0))
}
func rs() {
	defer wg.Done()
	for i := 0; i < len(lst); i += 1 {
		lst[i] = 555
	}
}
func rs2() {
	defer wg.Done()
	for i := 0; i < len(lst); i += 1 {
		lst[i] = 888
	}
}

// 并发的写map不支持
func ms() {
	defer wg.Done()
	for i := 0; i < 100; i += 1 {
		mp.Store(i, i)
	}
}
