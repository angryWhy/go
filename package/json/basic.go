package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age  int
		Addr string
	}
	p1 := Person{
		"taoge",
		30,
		"China", // oh my god, this comma cannot be omitted
	}
	data, _ := json.Marshal(p1)
	fmt.Println(string(data))
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() int {
	x := 5
	defer func(x int) {
		x++
	}(x)
	return 5
}
