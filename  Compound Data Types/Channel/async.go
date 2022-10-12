package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	go func() {
		time.Sleep(5 * time.Second)
		<-ch
		fmt.Println("子协程1--- 结束")
	}()
	//子协程1已启动，现在塞不进去，希望协程1执行，保证自己写入
	ch <- struct{}{}
	go func() {
		time.Sleep(5 * time.Second)
		//阻塞在这一行
		ch <- struct{}{}
	}()
	time.Sleep(10 * time.Second)
}
