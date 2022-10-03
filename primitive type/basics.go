package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	var b byte
	var c float32
	var d float64
	var f bool
	var s string
	var r rune
	var sli []int
	fmt.Printf("default value int = %d\n", a)
	fmt.Printf("default value byte = %d\n", b)
	fmt.Printf("default value float32 = %f\n", c)
	fmt.Printf("default value float64 = %f\n", d)
	fmt.Printf("default value bool = %t\n", f)
	fmt.Printf("default value string = %s\n", s)
	fmt.Printf("default value rune = %c\n", r)
	fmt.Printf("default value slcie = %v\n,arr is nil%t\n", sli, sli == nil)
}
func point() {
	var a int
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	//取址
	var ptr *int = &a
	fmt.Println(p, ptr)
}
