package main

import (
	"net/http"
)

func main() {
	http.Handle("/", new(foo).process(test))
	http.Handle("/test", new(foo).pipe(bar).process(test))
	http.Handle("/check", new(foo).auth().process(test))
	http.ListenAndServe(":8080", nil)
}

func foo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo("))
		next(w, r)
		w.Write([]byte(")"))
	}
}

func bar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bar("))
		next(w, r)
		w.Write([]byte(")"))
	}
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth("))
		next(w, r)
		w.Write([]byte(")"))
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

type middleware func(http.HandlerFunc) http.HandlerFunc

type pipeline struct {
	middlewares []middleware
}

func new(ms ...middleware) pipeline {
	return pipeline{append([]middleware(nil), ms...)}
}

func (p pipeline) pipe(ms ...middleware) pipeline {
	return pipeline{append(p.middlewares, ms...)}
}

func (p pipeline) auth(ms ...middleware) pipeline {
	return pipeline{append(p.middlewares, auth)}
}

func (p pipeline) process(h http.HandlerFunc) http.HandlerFunc {
	for i := range p.middlewares {
		h = p.middlewares[len(p.middlewares)-1-i](h)
	}
	return h
}