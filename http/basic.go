package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/boy", handleboy)
	//把服务起来,不发生err，就会一直阻塞
	http.ListenAndServe(":8080", nil)
}
func handleboy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	for k, v := range r.Header {
		fmt.Printf("%s = %v\n", k, v)
	}
	io.Copy(os.Stdout, r.Body)
	fmt.Fprintln(w, "hello")
}
