package main

import (
	jsoniter "github.com/json-iterator/go"
	http2 "go-code/study/http"
	"log"
)

func main() {
	//resp := http2.MyHttpPost2(http2.PostUrl, map[string]interface{}{"name": "mike", "age": 23})
	//resp := http2.MyHttpGet3("", nil)
	//resp := http2.MyFastHttpPost3("", map[string]interface{}{"name": "mike", "age": 23})
	//resp := http2.MyHystrixGet()
	resp := http2.MyHeimdallPost("", map[string]interface{}{"name": "mike", "age": 23})
	var r result
	err := jsoniter.Unmarshal([]byte(resp), &r)
	log.Println(err)
	log.Printf("%+v\n", r)
}

type result struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}
