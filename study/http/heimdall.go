package http

import (
	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/gojektech/heimdall/v6/hystrix"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	hc = httpclient.NewClient(httpclient.WithHTTPTimeout(time.Second*5), httpclient.WithRetryCount(3))
)

type myHttpClient struct {
	client http.Client
}

func (c *myHttpClient) Do(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth("username", "password")
	request.Header.Add("token", "xxx")
	return c.client.Do(request)
}

func MyHeimdallGet() string {
	//client := httpclient.NewClient(httpclient.WithHTTPTimeout(time.Second*5), httpclient.WithRetryCount(3))
	//client := httpclient.NewClient(httpclient.WithHTTPClient(&myHttpClient{
	//	client: http.Client{Timeout: time.Second*5},
	//}))
	resp, err := hc.Get(GetUrl, nil)
	if err != nil {
		log.Println("heimdall get error:", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("heimdall read error:", err)
		return ""
	}
	log.Println(string(body))
	return string(body)
}

func MyHeimdallGet2() string {
	req, err := http.NewRequest(http.MethodGet, GetUrl, nil)
	if err != nil {
		log.Println("heimdall NewRequest error:", err)
		return ""
	}
	resp, err := hc.Do(req)
	if err != nil {
		log.Println("heimdall Do error:", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("heimdall read error:", err)
		return ""
	}
	log.Println(string(body))
	return string(body)
}

func MyHystrixGet() string {
	fallbackFn := func(err error) error {
		//_, err := http.Post(PostUrl)
		return err
	}

	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(100 * time.Millisecond),
		hystrix.WithCommandName("my_hystrix_get"),
		hystrix.WithHystrixTimeout(1000 * time.Millisecond),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
		hystrix.WithStatsDCollector("localhost:8125", "myapp.hystrix"),
		hystrix.WithFallbackFunc(fallbackFn),
	)

	resp, err := client.Get(GetUrl, nil)
	if err != nil {
		log.Println("heimdall get error:", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("heimdall read error:", err)
		return ""
	}
	log.Println(string(body))
	return string(body)
}

func MyHeimdallPost(u string, params map[string]interface{}) string {
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Add(k, cast.ToString(v))
	}

	req, err := http.NewRequest(http.MethodPost, PostUrl, strings.NewReader(urlValues.Encode()))
	if err != nil {
		log.Println("heimdall NewRequest error:", err)
		return ""
	}
	req.Header.Add("token", "ttt")
	req.Header.Add("Content-Type", "application/json")

	resp, err := hc.Do(req)
	if err != nil {
		log.Println("heimdall Do error:", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("heimdall read error:", err)
		return ""
	}
	log.Println(string(body))
	return string(body)
}
