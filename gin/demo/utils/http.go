package utils

import (
	"github.com/spf13/cast"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

const (
	HttpTimeout = 5
)

func HttpGet(url string) (ret []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	if err = fasthttp.DoTimeout(req, resp, time.Second*HttpTimeout); err != nil {
		Log.Errorf("fasthttp HttpGet error: %s", err.Error())
	}

	ret = resp.Body()
	Log.Infof("HttpGet resp: %s", string(ret))
	return
}

func HttpGetWithParam(url string, params map[string]interface{}) (ret []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url + "?" + HttpQueryBuild(params))
	resp := fasthttp.AcquireResponse()
	if err = fasthttp.DoTimeout(req, resp, time.Second*HttpTimeout); err != nil {
		Log.Errorf("fasthttp HttpGetWithParam error: %s", err.Error())
	}

	ret = resp.Body()
	Log.Infof("HttpGetWithParam resp: %s", string(ret))
	return
}

func HttpPost(url string, params map[string]interface{}) (ret []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	resp := fasthttp.AcquireResponse()
	req.SetBody([]byte(HttpQueryBuild(params)))

	if err = fasthttp.DoTimeout(req, resp, time.Second*HttpTimeout); err != nil {
		Log.Errorf("fasthttp HttpPost error: %s", err.Error())
	}

	ret = resp.Body()
	Log.Infof("httpPost resp: %s", string(ret))
	return
}

func HttpPostJson(url string, json []byte) (ret []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetBody(json)
	resp := fasthttp.AcquireResponse()

	if err = fasthttp.DoTimeout(req, resp, time.Second*HttpTimeout); err != nil {
		Log.Errorf("fasthttp HttpPostJson error: %s", err.Error())
	}

	ret = resp.Body()
	Log.Infof("httpPostJson resp: %s", string(ret))
	return
}

func HttpQueryBuild(params map[string]interface{}) string {
	q := make(url.Values)
	for k, v := range params {
		q.Add(k, cast.ToString(v))
	}
	return q.Encode()
}
