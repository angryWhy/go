package main

import "fmt"

func main() {
	makeSlice()
}
func makeSlice() {
	arr := [...]string{"haha", "hee", "kk", "ll", "aaa"}
	//引用全部数组
	all := arr[:]
	//下标1到末尾
	part := arr[1:]
	//引用3，4个元素
	part2 := arr[3:5]
	fmt.Printf("%v,%v,%v", all, part, part2)
}
