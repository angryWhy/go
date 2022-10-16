package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	countCh := make(chan int)
	finishChan := make(chan struct{})
	go count(countCh, 10, finishChan)
	abortCh := make(chan struct{})
	go abort(abortCh)
LOOP:
	for {

		select {
		case n := <-countCh:
			fmt.Println(n)
		case <-finishChan:
			fmt.Println("finish")
			break LOOP
		case <-abortCh:
			fmt.Println("abort")
			break LOOP
		}
	}
}
func count(countCh chan int, n int, finishChan chan struct{}) {
	if n < 0 {
		return
	}
	ticker := time.NewTicker(1 * time.Second)
	for {
		countCh <- n
		<-ticker.C
		n--
		if n <= 0 {
			ticker.Stop()
			finishChan <- struct{}{}
			break
		}
	}

}

// 终止
func abort(ch chan struct{}) {
	buffer := make([]byte, 1)
	os.Stdin.Read(buffer)
	ch <- struct{}{}
}
