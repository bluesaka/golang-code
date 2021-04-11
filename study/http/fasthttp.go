package http

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

var (
	fc = &fasthttp.Client{}
)

func MyFastHttpGet() string {
	status, body, err := fasthttp.Get(nil , GetUrl)
	if err != nil {
		log.Println("fasthttp Get error:", err)
		return ""
	}

	if status != fasthttp.StatusOK {
		log.Println("fasthttp Get status error:", status)
		return ""
	}

	log.Println(string(body))
	return string(body)
}

func MyFastHttpPost(u string, params map[string]interface{}) string {
	args := &fasthttp.Args{}
	for k, v := range params {
		args.Add(k, cast.ToString(v))
	}
	status, body, err := fasthttp.Post(nil , PostUrl, args)
	if err != nil {
		log.Println("fasthttp Post error:", err)
		return ""
	}

	if status != fasthttp.StatusOK {
		log.Println("fasthttp Post status error:", status)
		return ""
	}

	log.Println(string(body))
	return string(body)
}

func MyFastHttpPost2(u string, params map[string]interface{}) string {
	req := &fasthttp.Request{}
	req.SetRequestURI(PostUrl)

	//reqBody := []byte(`{"name":"james"}`)
	reqBody, _ := jsoniter.Marshal(params)
	log.Println(string(reqBody))
	//for k, v := range params {
	//	reqBody = append(reqBody)
	//}
	req.SetBody(reqBody)

	// 默认是 application/x-www-form-urlencoded
	req.Header.SetContentType("Application/json")
	req.Header.SetMethod("POST")
	req.Header.Add("token", "xxx")
	req.Header.Add("x-header", "abc")

	resp := &fasthttp.Response{}
	if err := fc.Do(req, resp); err != nil {
		log.Println("fasthttp Post error:", err)
		return ""
	}

	body := resp.Body()
	log.Println(string(body))
	return string(body)
}

// AcquireRequest 性能更高
func MyFastHttpPost3(u string, params map[string]interface{}) string {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(PostUrl)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	reqBody, _ := jsoniter.Marshal(params)
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	//fasthttp.Do(req, resp)
	// Do with Timeout
	if err := fasthttp.DoTimeout(req, resp, time.Second*5); err != nil {
		log.Println("fasthttp Post error:", err)
		return ""
	}

	body := resp.Body()
	log.Println(string(body))
	return string(body)
}

func PrintFastHttpClient() {
	log.Printf("%+v\n", fc)
}
