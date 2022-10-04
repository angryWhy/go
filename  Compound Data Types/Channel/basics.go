package main

import "fmt"

func main() {
	var ch chan int
	fmt.Printf("ch is nil---%t\n", ch == nil)
	fmt.Printf("ch is len---%d\n", len(ch))
	ch2 := make(chan int, 8)
	fmt.Printf("ch2 len%d,cap%d\n", len(ch2), cap(ch2))
	for i := 0; i < 10; i++ {
		ch2 <- 9
	}
	fmt.Printf("ch2 len%d,cap%d\n", len(ch2), cap(ch2))
}
