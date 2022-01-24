package main

import (
	"io"
	"os"
)

func main() {
	mystring := "hello world"
	io.WriteString(os.Stdout, mystring)
}
