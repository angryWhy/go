package main

import "fmt"

func main() {
	mul()
}
func basic() {
	var a = 3
me:
	a += 3
	goto me
}
func mul() {
	i := 4
	if i%2 == 0 {
		goto l1
	} else {
		goto l2
	}
l1:
	fmt.Printf("这是l1\n")
	i += 3
	fmt.Printf("i的值%d\n", i)
l2:
	i *= 7
	fmt.Printf("这是l2\n")
	fmt.Printf("i的值%d\n", i)
}
