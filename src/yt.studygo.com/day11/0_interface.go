package main

import "io"

// 实现跟strings.NewReader类似的功能，用来解析html
// strings.NewReader() 方法返回的是一个 Reader对象，Reader对象实现了io.Reader()接口
type htmlParser struct {
	html string
	i    int64
}

func (hps *htmlParser) Read(b []byte) (n int, err error) {
	// 实现reader接口 将html数据写入到b中
	if hps.i > int64(len(hps.html)) {
		return 0, io.EOF
	}
	n = copy(b, hps.html[hps.i:])
	hps.i += int64(n)
	return
}

func (hps *htmlParser) Write(b []byte) (n int, err error) {
	// 实现writer接口 将html的内容输出
	return
}

//接收字符串输入的HTML解析器
func NewHtmlParser(html string) *htmlParser {
	//解析html参数构建html
	return &htmlParser{html, 0}
}

/* func main() {
	//htmlmap := make(map[string]string, 1)
	//htmlmap["h1"] = "hello, world"
	html := "<h1>hello, world</h1>"
	NewHtmlParser(html)
}
*/
