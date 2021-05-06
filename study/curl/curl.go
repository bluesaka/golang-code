package curl

/**
curl  (client url工具)

@link http://www.ruanyifeng.com/blog/2019/09/curl-reference.html

// query 数据格式
curl 'http://httpbin.org/get?a=1&b=2'

// `curl -d` 发起post请求，等同于 `curl -d -X POST`
// 请求参数默认 Content-Type 为 x-www-form-urlencoded (Form表单)
curl -d 'a=1&b=2' -d 'c=3' http://httpbin.org/post

// 设置请求头 JSON格式
curl -d '{"aa":11, "bb": 22}' -H 'Content-Type: application/json' -H 'Token: xxx' 'http://httpbin.org/post'

 */