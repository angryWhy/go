package main

import "fmt"

type sss int8

func main() {
	var a sss
	fmt.Println(a)
}

type ms map[int]int

func (m ms) add(a int) {
	fmt.Printf("%d", m[0])
}
