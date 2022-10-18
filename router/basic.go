package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	//router.GET("/", get)
	router.GET("/user/:name/:age", getParams)
	router.POST("/", post)
	//返回静态文件
	router.ServeFiles("/file/*filepath", http.Dir("./static"))
	http.ListenAndServe(":8080", router)
}
func get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("request method:%s\n", r.Method)
	fmt.Printf("request body :")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
	//
	fmt.Fprint(w, "hello get")
	//等价
	//w.Write([]byte("hello get"))
}
func post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("request method:%s\n", r.Method)
	fmt.Printf("request body :")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
	//
	fmt.Fprint(w, "hello post")
	//等价
	w.Write([]byte("hello post"))
}
func getParams(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("name=%s age=%s \n", params.ByName("name"), params.ByName("age"))
}
