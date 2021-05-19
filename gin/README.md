### start

```
go build main.go

nohup ./main > /dev/null &

平滑重启：
kill -USR2 {pid}
```