package main

import (
	"fmt"
	"io"
	"net/http"
)

/* func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "hello, %s!", name)
} */

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

// 重构
func Greet(writer io.Writer, name string) { // 因为之前的参数都实现了io.Writer接口
	fmt.Fprintf(writer, "hello, %s!", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "chirs")
}
