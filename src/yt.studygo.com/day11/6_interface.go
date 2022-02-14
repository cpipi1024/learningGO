package main

/* type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("%.2f$", d) // 返回根据format生成的字符串
}

// 实现handler接口
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price) // 将根据format生成的字符串写入到w中
	}
}
*/
/* func main() {
	db := database{"shoes": 50, "shirts": 100}
	log.Fatal(http.ListenAndServe("localhost:8000", db)) //监听请求 并且转发给handler处理
} */
