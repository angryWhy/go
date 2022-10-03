package main

import (
	"fmt"
	"strings"
)

func main() {
	join_str()
}
func join_str() {
	s1 := "a"
	s2 := "b"
	s3 := "c"

	str := s1 + " " + s2 + " " + s3
	str2 := fmt.Sprintf("%s %s %s", s1, s2, s3)
	str3 := strings.Join([]string{s1, s2, s3}, " ")
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	s4 := sb.String()
	fmt.Println(str, str2, str3, s4)
}
