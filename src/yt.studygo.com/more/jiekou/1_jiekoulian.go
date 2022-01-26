package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//  使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
type wordcounter struct {
	words int
	lines int
}

//wordcounter实现了Read()签名函数 可以看作io.Reader抽象接口的实例

func (wc *wordcounter) Read(b []byte) (n int, err error) {
	wc.words = len(strings.Split(string(b), ","))
	return 0, nil
}

func readfile() {
	var count int = 0
	var line int = 0
	file, err := os.Open("danci.txt")
	if err != nil {
		panic("!!!ERROR!!!")
	}
	new_scanner1 := bufio.NewScanner(file)
	new_scanner1.Split(bufio.ScanWords)
	for new_scanner1.Scan() {
		count++
		text := new_scanner1.Text()
		fmt.Printf("text: %s\n", text)
	}
	new_scanner2 := bufio.NewScanner(file)
	new_scanner2.Split(bufio.ScanLines)
	for new_scanner2.Scan() {
		line++
		text := new_scanner2.Text()
		fmt.Printf("text: %s\n", text)
	}
	fmt.Printf("count: %d, line: %d\n", count, line)
}

func main() {
	/* wc := wordcounter{}
	wc.Read([]byte("nihao,hello,haha"))
	fmt.Println(wc.words) */
	readfile()
}
