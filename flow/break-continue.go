package main

import "fmt"

func main() {

}
func continuefn() {
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		if i > 3 {
			continue
			//直接进入下一层循环
		}
		fmt.Printf("%d", v)
	}
}
func breakfn() {
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		if i > 3 {
			fmt.Printf("%d", v)
		}
	}
}
