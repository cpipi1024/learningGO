package main

import (
	"io"
)

// 不包含任何值的nil接口值和刚好包含nil值的接口值是完全不同的

const debug = true

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("hello world \n"))
	}
}

/* func main() {
	var buff *bytes.Buffer
	if debug {
		buff = new(bytes.Buffer)
	}
	f(buff)
	if debug {
		fmt.Fprintf(buff, "none")
	}
}
*/
