package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	names := map[int]int{1: 2}
	fmt.Println(ages, names)
	var mapa map[int]int
	mapa[1] = 1
	fmt.Println(mapa == nil)
}
func sortNames(names map[int]int) {
	sl := make([]int, 0, len(names))
	for _, v := range names {
		sl = append(sl, v)
	}
	sort.Ints(sl)
	for _, v := range sl {
		fmt.Println(v)
	}
}
func define() {
	names := map[int]int{1: 2}
	value, ok := names[1]
	if ok {
		//执行操作
	}
	if value, ok := names[1]; !ok {
		//执行操作操作
	}
}
func equal(x, y map[int]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if value, ok := y[k]; !ok || value != v {
			return false
		}
	}
	return true
}
