package main

import "fmt"

func main() {
	color := "black"
	switch color {
	case "black":
		fmt.Println()
	case "red":
		fmt.Println()
	default:
		fmt.Println()
	}
}
func switch_type() {
	var num interface{} = 6.5
	switch value := num.(type) {
	case int:
		fmt.Printf("int类型-%d", value)
	case string:
		fmt.Printf("string类型-%s", value)
	}
}
