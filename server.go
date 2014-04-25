package main

import (
	"io"
	"net/http"
	"runtime"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", "text/plain",
	)
	io.WriteString(res, "hello there!")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
