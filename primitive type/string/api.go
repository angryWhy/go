package main

import (
	"fmt"
	"strings"
)

func main() {
	contains_str()
}

// 分割
func split_str() {
	s := "我是,123"
	arr := strings.Split(s, ",")
	fmt.Println(arr, arr[0])
}

// 包含
func contains_str() {
	s := "hello world"
	fmt.Printf("%t\n", strings.Contains(s, "e"))
}

// 开头结尾
func before_after() {
	s := "hello world"
	fmt.Printf("%t\n", strings.HasPrefix(s, "he"))
	fmt.Printf("%t\n", strings.HasSuffix(s, "ld"))
}
