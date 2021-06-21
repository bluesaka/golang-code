package main

import (
	"bytes"
	"encoding/json"
	json2 "go-code/study/json"
	"log"
	"strings"
)

func main() {
	u := json2.User{
		Name: "hello & world",
		Age:  18,
	}
	b, _ := json2.MyMarshal(u)
	log.Println(b)

	user := json2.MyUnmarshal(b)
	log.Printf("%+v\n", user)

	// json Marshal `hello & world` -> `hello \u0026 world`
	b2, _ := json.Marshal(u)
	log.Println(string(b2))

	m := make([]string, 0, 1)
	m = append(m, "hello & world")
	b3, _ := json.Marshal(m)
	log.Println(b3)
	log.Println(string(b3))
	s1 := strings.Replace(string(b3), "\\u0026", "&", -1)
	b3 = bytes.Replace(b3, []byte("\\u0026"), []byte("&"), -1)
	log.Println(b3)
	log.Println(string(b3))
	log.Println(s1)

	NewJsonEncoder()
}

/**
json.Marshal 默认 escapeHtml 为true, 会转义 <、>、&

解决方法：
1. replace

b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)

s = strings.Replace(s, "\\u0026", "&", -1)
s = strings.Replace(s, "\\u003c", "<", -1)
s = strings.Replace(s, "\\u003e", ">", -1)

2. New json encoder
*/
func NewJsonEncoder() {
	buf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode("hello & world"); err != nil {
		log.Fatal(err)
	}
	log.Println(buf.String())
}
