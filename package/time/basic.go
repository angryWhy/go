package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := "2022-10-10 10:10:10"
	now := time.Now()
	ts := now.Format(time1)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(time1, ts, loc)
	fmt.Println(t)
	timesub()
}
func timesub() {
	time1 := "2022-10-10 10:10:10"
	now := time.Now()
	fmt.Println(now.Day(), now.Weekday(), now.Year(), time1)
	fmt.Println(now.Unix())
}
func tk() {
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-tk.C
	}
	tk.Stop()
}
func timer() {
	ti := time.NewTimer(1 * time.Second)
	<-ti.C
	ti.Stop()
}
