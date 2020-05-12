多路复用器

URL 匹配

```
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

type greetingHandler struct {
	Name string
}

func (h greetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", h.Name)
}

func main() {
	mux := http.NewServeMux()
	// 注册处理器函数
	mux.HandleFunc("/hello", helloHandler)

	// 注册处理器
	mux.Handle("/greeting/golang", greetingHandler{Name: "Golang"})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
```
