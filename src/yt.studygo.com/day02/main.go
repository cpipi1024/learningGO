package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var who string
	who = "Hello "
	if len(os.Args) > 1 {
		args := os.Args[1:]
		fmt.Printf("the type is %T:\n", args)
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(who)

}
