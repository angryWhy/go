package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	var n int32
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			n++
			//atomic.AddInt32(&n, 1)
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
