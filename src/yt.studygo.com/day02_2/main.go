package main

import (
	"fmt"
)

const englishPreFix = "Hello, "

func hello(name string) string {
	return englishPreFix + name
}

func main() {
	fmt.Println(hello("chris"))
}
