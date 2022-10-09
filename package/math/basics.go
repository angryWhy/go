package main

import (
	"fmt"
	"math/rand"
)

func main() {
	source := rand.NewSource(20)
	rander := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Println(rander.Intn(100))
	}
	rander2 := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Println(rander2.Intn(100))
	}
}
