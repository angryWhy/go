package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	get()
	post()
	complexHttpRequest()
}
func get() {
	resp, err := http.Get("http://localhost:8080/boy")
	if err != nil {
		panic(err)
	}
	for k, v := range resp.Header {
		fmt.Printf("%s = %v", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
func post() {
	reader := strings.NewReader("hello serve")
	resp, err := http.Post("http://localhost:8080/boy", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	for k, v := range resp.Header {
		fmt.Printf("%s = %v", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
func complexHttpRequest() {
	//请求体，io.reader
	reader := strings.NewReader("hello complex")
	re, err := http.NewRequest("POST", "http://localhost:8080/boy", reader)
	if err != nil {
		panic(err)
	}
	//自定义请求头
	re.Header.Add("User-Agent", "china")
	re.Header.Add("User-Agent", "head-value")
	//自定义cookie
	re.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: "haha",
	})
	//创建cilent
	cilent := &http.Client{}
	//提交http请求
	cilent.Do(re)
	defer re.Body.Close()
}
