### APISIX
> APISIX是apache旗下，基于Openresty + etcd的云原生、高性能、可扩展的微服务API网关。
> 通过插件机制，提供了动态负载平衡、身份验证、限流限速等功能，也可自行开发插件进行拓展。


#### openresty
```
http://openresty.org/cn/download.html

./configure --prefix=/usr/local/src/openresty \
--with-cc-opt="-I/usr/local/include -I/usr/local/Cellar/openssl@1.1/1.1.1h/include -I/usr/local/Cellar/pcre/8.44/include" \
--with-ld-opt="-L/usr/local/Cellar/openssl@1.1/1.1.1h/lib/ -L/usr/local/Cellar/pcre/8.44/lib/"
-j8

make -j8
make install
```

#### LuaRocks
```
brew install luarocks lua@5.1
```

#### apisix
```
ln -s /usr/local/Cellar/lua@5.1/5.1.5_8 /opt/lua@5.1
ln -s /usr/local/Cellar/openssl@1.1/1.1.1h /usr/local/src/openresty/openssl

make deps

```


#### 遇到的问题
```
问题：'http_stub_status_module' module is missing in your openresty, please check it out. Without this module, there will be fewer monitoring indicators.
nginx: [emerg] unknown directive "real_ip_header" in /usr/local/src/apache-apisix-2.7/conf/nginx.conf:109

解决：重新编译openresty
./configure --prefix=/usr/local/src/openresty \
--with-http_stub_status_module \
--with-http_realip_module \
--with-http_v2_module \
--with-cc-opt="-I/usr/local/include -I/usr/local/Cellar/openssl@1.1/1.1.1h/include -I/usr/local/Cellar/pcre/8.44/include" \
--with-ld-opt="-L/usr/local/Cellar/openssl@1.1/1.1.1h/lib/ -L/usr/local/Cellar/pcre/8.44/lib/"
```

#### 使用
```
apisix默认端口 9080
curl http://localhost:9080

版本：./bin/apisix version
启动：./bin/apisix start
重启：./bin/apisix restart
停止：./bin/apisix stop
```

#### 创建router
```
curl http://127.0.0.1:9080/apisix/admin/routes/1 -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X PUT -d '
{
    "upstream": {
       "nodes": {
           "httpbin.org:80": 1
       },
       "type": "roundrobin"
    },
    "uri": "/get"
}'
```

#### 访问router
```
curl -i http://127.0.0.1:9080/get
```