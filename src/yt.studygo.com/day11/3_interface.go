package main

// 7.5 接口的值
/*
func main() {
	var w io.Writer
	if w == nil {
		fmt.Printf("w 的接口值是nil\n")
	}
	var b *bytes.Buffer
	if b == nil {
		fmt.Printf("b 的接口值是nil\n")
	}
	// 以上两种接口的接口值在创建并且初始化时被确定为nil
	// 同样的在调用接口的方法时会panic
	//w.Write([]byte("hello")) // panic

	w = os.Stdout // 调用了具体类型到接口类型的隐式转换 相当于io.Writer(os.Stdout)

	w.Write([]byte("hello world"))

}
*/
