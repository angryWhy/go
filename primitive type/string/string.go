package main

import "fmt"

func main() {
	str := "我是123"
	fmt.Printf("字符串长度%d\n", len(str))
	for _, ele := range str {
		fmt.Println(ele)
	}

	arr := []byte(str)
	for _, ele := range arr {
		fmt.Println(ele)
	}
	fmt.Printf("字符串byte数组的长度%d\n", len(arr))
	arr2 := []rune(str)
	for _, ele := range arr2 {
		fmt.Println(ele)
	}
	fmt.Printf("字符串rune数组的长度%d\n", len(arr2))
}
