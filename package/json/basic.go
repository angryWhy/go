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
}
