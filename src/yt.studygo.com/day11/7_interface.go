package main

import (
	"fmt"
	"net/http"
)

// 7.8 http.handler接口
// 使用http.mux 注册多个 handler
type dollars float32
type database map[string]dollars

// 实现Stringer接口
func (d dollars) String() string {
	return fmt.Sprintf("%.2f$", d)
}

func (db database) list(rw http.ResponseWriter, req *http.Request) {
	// 访问/list的handler
	for item, price := range db {
		fmt.Fprintf(rw, "%s:%s\n", item, price)
	}
}

func (db database) price(rw http.ResponseWriter, req *http.Request) {
	// 访问/price的handler
	item := req.URL.Query().Get("item")
	prices, ok := db[item]
	if !ok {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "not found item: %s\n", item)
		return
	}
	fmt.Fprintf(rw, "%s:%s\n", item, prices)
}
func main() {
	// database
	db := database{"shoes": 50.0, "shirts": 100.0, "hat": 10.0}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list)) // http.HandlerFunc是一个适配器 通过将具有相同签名的函数转换为HandlerFunc类型
	mux.Handle("/price", http.HandlerFunc(db.price))
	http.ListenAndServe("localhost:8000", mux)
}
