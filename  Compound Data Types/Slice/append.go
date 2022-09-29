package main

import "fmt"

func main() {
	var runes []rune
	for _, v := range "hello,世界" {
		runes = append(runes, v)
	}
	fmt.Printf("%q\n", runes)
	s := make([]int, 3)
	fmt.Printf("原切片%v,%d,%d，%p\n", s, len(s), cap(s), &s)
	appendSlice(s)
	fmt.Printf("原切片%v,%d,%d,%p\n", s, len(s), cap(s), &s)
	useAppend()
}
func appendSlice(s []int) {
	s[0] = 9
	s = append(s, 1)
	//fmt.Printf("%v", s)
	fmt.Printf("子切片%v,%d,%d,%p\n", s, len(s), cap(s), &s)
}
func useAppend() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data) //"one","three"
	res := noempty1(data)
	fmt.Printf("%q\n", res)  //"one","three"
	fmt.Printf("%q\n", data) //"one", "three", "three"
}
func noempty1(strings []string) []string {
	//相当于应用原始的slice的新的零长度的slice
	out := strings[:0]
	for _, v := range strings {
		if v != "" {
			out = append(out, v)
		}
	}
	return strings
}
