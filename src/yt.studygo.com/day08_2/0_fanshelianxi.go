package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//使用反射机制完成读取配置文件和按照配置文件中内容调用函数
type Person struct {
	Name  string `ini:"name"`
	Age   int    `ini:"age"`
	Place string `ini:"place"`
}

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

func strToMap(str string) map[string]string {
	// 分割字符串
	var strslice1 []string
	var strslice2 []string
	var strmap map[string]string
	strmap = map[string]string{}

	strslice1 = strings.Split(str, ";")
	for i := 0; i < len(strslice1); i++ {
		strslice2 = append(strslice2, strings.Split(strslice1[i], "=")...)
	}
	fmt.Printf("分割字符串后的结果: %v\n", strslice2)
	for i := 0; i < len(strslice2); i += 2 {
		strmap[strslice2[i]] = strslice2[i+1]
	}
	return strmap
}

func reflectMap(a interface{}, m map[string]string) {
	k := reflect.TypeOf(a)  // Type类型
	v := reflect.ValueOf(a) // Value类型

	fmt.Printf("type: %v, type.Name() : %v, type.Kind(): %v\n", k, k.Name(), k.Kind())
	fmt.Printf("value.Type(): %v, value.Type().Name(): %v, value.kind(): %v\n", v.Type(), v.Type().Name(), v.Kind())
	fmt.Printf("type.Elem(): %v\n", k.Elem()) //
	fmt.Printf("value.Elem(): %v\n", v.Elem())
	//kind()返回底层类型reflect.XXX
	for i := 0; i < v.Elem().NumField(); i++ { // type.NumField() 返回struct类型的字段数
		// 根据m中的键判断类型
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

/* func invokeFunc() {
} */

func main() {
	str := orFile("canshu.ini")
	strmap := strToMap(str)
	p := Person{}
	fmt.Printf("配置文件映射: %v\n", strmap)
	reflectMap(&p, strmap)
	fmt.Printf("p: %v", p)
}
