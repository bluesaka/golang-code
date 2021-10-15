package main

import (
	"fmt"
	"net/http"
)

func main() {
	server2()
}

func server1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", mux)
}

func server2() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	w.Write([]byte("hello world"))
}