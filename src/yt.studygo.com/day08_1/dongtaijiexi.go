package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 利用反射学到的知识动态解析ini配置文件

func reflecValue(content string) {
	split := strings.Split(content, ";")
	fmt.Println("结果", split)
	for _, v := range split {
		switch {
		case strings.Contains(v, "name"):
			fmt.Println("name is stirng")
		case strings.Contains(v, "age"):
			fmt.Println("age is integer")
		default:
			break
		}
	}
}
func readFile(f *os.File) {
	// 读取文件内容
	content := []byte{}
	tmp := make([]byte, 128)
	for {
		n, err := f.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件报错")
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Printf("读取文件内容:\n%s\n", string(content[:]))
	reflecValue(string(content))
}

func main() {
	file, _ := os.Open("peizhi.ini")
	readFile(file)
}
