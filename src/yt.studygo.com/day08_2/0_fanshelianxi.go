package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//使用反射机制完成读取配置文件和按照配置文件中内容调用函数

func orFile(filepath string) string {
	var line []byte
	var isprefix bool
	var err error
	var final_line []byte
	file, err := os.Open(filepath)
	if err != nil {
		panic("读取文件失败！")
	}
	reader := bufio.NewReader(file)

	//line, isprefix, err = reader.ReadLine()
	for {
		line, isprefix, err = reader.ReadLine()
		final_line = append(final_line, line...)
		if err == io.EOF { //读到结尾时会返回io.EOF
			break
		}

	}
	fmt.Printf("读取到的数据: %v, 返回标志: %t, 返回的错误信息: %v\n", string(final_line), isprefix, err)
	str_line := string(final_line)
	return str_line
}

func strToMap(str string) map[string]interface{} {
	// 分割字符串
	var strslice1 []string
	var strslice2 []string
	var strmap map[string]interface{}

	strslice1 = strings.Split(str, ";")
	for i := 0; i < len(strslice1); i++ {
		strslice2 = append(strslice2, strings.Split(strslice1[i], "=")...)
	}
	fmt.Printf("分割字符串后的结果: %v\n", strslice2)

	return strmap
}

/* func reflectFille() map[string]interface{} {

} */

/* func invokeFunc() {
} */

func main() {
	str := orFile("canshu.ini")
	strmap := strToMap(str)
	fmt.Printf("配置文件映射: %v", strmap)
}
