package main

import "fmt"

func main() {
	makeSlice()
	//output:5,4,3,2,1
	reverse(arr[:])
	fmt.Printf("%v\n", arr)
	slice_init()
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

var arr = [...]int{1, 2, 3, 4, 5}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func demo() {
	data := []string{"one", "", "three"}
	res := noempty(data)
	fmt.Printf("%q\n", res)  //"one","three"
	fmt.Printf("%q\n", data) //"one", "three", "three"

}

func noempty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings
}
func slice_init() {
	s2 := [][]int{
		{1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Printf("%d,%d,%d,%d", len(s2), len(s2[0]), len(s2[1]), len(s2[2]))
}
