package main

// 实现跟strings.NewReader类似的功能，用来解析html
// strings.NewReader() 方法返回的是一个 Reader对象，Reader对象实现了io.Reader()接口
type htmlParser struct {
}

func (hps *htmlParser) Read(b []byte) (int, error) {

}

func NewHtmlParser(html string) *htmlParser {
	return
}

func main() {
	html := "<h1>hello, world</h1>"

}
