package http

import (
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	c = &http.Client{
		Timeout: 5 * time.Second,
	}

	GetUrl  = "http://httpbin.org/get"
	PostUrl = "http://httpbin.org/post"
)

func MyHttpGet() string {
	resp, err := http.Get(GetUrl)
	if err != nil {
		log.Println("HTTP Get error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Get body error:", err)
		return ""
	}
	log.Println(string(body))
	log.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println("HTTP Get success")
	}

	return string(body)
}

func MyHttpGet2(u string, params map[string]interface{}) string {
	q := url.Values{}
	//q := make(url.Values)
	//q.Set("name", "alice")
	//q.Set("age", "22")
	for k, v := range params {
		q.Set(k, cast.ToString(v))
	}
	path, _ := url.Parse(u)
	path.RawQuery = q.Encode()
	log.Println(path.String())

	resp, err := http.Get(path.String())
	if err != nil {
		log.Println("HTTP Get error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Get body error:", err)
		return ""
	}
	log.Println(string(body))
	log.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println("HTTP Get success")
	}

	return string(body)
}

func MyHttpGet3(u string, params map[string]interface{}) string {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Println("HTTP NewRequest Get error:", err)
		return ""
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Set(k, cast.ToString(v))
	}
	req.URL.RawQuery = q.Encode()

	// add header
	req.Header.Add("token", "xxx")
	req.Header.Add("x-header", "test-header")

	resp, err := c.Do(req)
	if err != nil {
		log.Println("HTTP Get error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Get body error:", err)
		return ""
	}
	log.Println(string(body))
	log.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println("HTTP Get success")
	}

	return string(body)
}

func MyHttpPost(u string, params map[string]interface{}) string {
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Add(k, cast.ToString(v))
	}
	resp, err := http.PostForm(u, urlValues)
	if err != nil {
		log.Println("HTTP Post error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Post body error:", err)
		return ""
	}
	log.Println(string(body))
	log.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println("HTTP Post success")
	}
	return string(body)
}

func MyHttpPost2(u string, params map[string]interface{}) string {
	q := url.Values{}
	for k, v := range params {
		q.Add(k, cast.ToString(v))
	}
	req, err := http.NewRequest("POST", u, strings.NewReader(q.Encode()))
	if err != nil {
		log.Println("HTTP NewRequest Post error:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		log.Println("HTTP Post error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Post body error:", err)
		return ""
	}
	log.Println(string(body))
	log.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println("HTTP Post success")
	}
	return string(body)
}
