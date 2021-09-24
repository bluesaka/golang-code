### start

```
go build main.go

nohup ./main > /dev/null &

平滑重启：
kill -USR2 {pid}

go build main.go && kill -USR2 `ps -ef | grep ./main | grep -v grep | awk '{print $2}'`
```