package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 利用反射学到的知识动态解析ini配置文件
type Person struct {
	Name  string `ini:"name"`
	Age   int    `ini:"age"`
	Phone string `ini:"phone"`
}

// a--结构体 m--文本内容字典
func reflecValue(a interface{}, m map[string]string) {
	/* split := strings.Split(content, ";")
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
	} */
	k := reflect.TypeOf(a)                     // 返回Type类型
	v := reflect.ValueOf(a)                    // 返回Value类型
	for i := 0; i < v.Elem().NumField(); i++ { // value.Elem()返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
		if tagvalue, ok := m[k.Elem().Field(i).Tag.Get("ini")]; ok {
			switch k.Elem().Field(i).Type.Kind() {
			case reflect.Int:
				num, _ := strconv.Atoi(tagvalue)
				v.Elem().Field(i).SetInt(int64(num))
			case reflect.String:
				v.Elem().Field(i).SetString(tagvalue)
			}
		}
	}

}
func readFile(f *os.File) map[string]string {
	// 读取文件内容
	var m = map[string]string{}
	tmp := make([]byte, 128)
	for {
		_, err := f.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件报错")
			break
		}
		str_tmp := string(tmp[:])
		split_tmp := strings.Fields(str_tmp) //返回将字符串按照空白字符分割的多个字符串，如果字符串全部是空白或者是空字符串的话，就会返回空切片。（注意空白字符！！！）
		//fmt.Println("split_tmp: ", len(split_tmp))
		for _, s := range split_tmp {
			var kv = strings.Split(s, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}
	return m
}

func main() {
	var m map[string]string
	file, _ := os.Open("peizhi.ini")
	m = readFile(file)
	var s = Person{}
	reflecValue(&s, m)
	fmt.Println(s)

}
