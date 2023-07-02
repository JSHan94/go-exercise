package main

import (
	// "net/http"

	hello "github.com/JSHan94/go-exercise/03-package/02-package/library"
	// "github.com/gorilla/mux"
)

func main() {
	hello.SayHello()
	hello.SayBye()

	// r := mux.NewRouter()
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, world!\n"))
	// })
	// http.ListenAndServe(":8080", r)
}
