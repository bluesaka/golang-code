package time

import (
	"log"
	"time"
)

func GetCurrDatetime() string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	//location, _ := time.LoadLocation("Local")
	return time.Now().In(location).Format("2006-01-02 15:04:05")
}

func GetCurrTimestamp() int64 {
	return time.Now().Unix()
}

func GetTimestamp(day string) int64 {
	location, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", day, location)
	return t.Unix()
}

// GetDatetime get date time
// date 2020-01-01
func GetDatetime(day string) time.Time {
	location, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", day, location)
	log.Println(t)
	return t
}

// GetDatetime2 get date time
// date 2020-01-01 01:01:01
func GetDatetime2(day string) time.Time {
	location, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", day, location)
	log.Println(t)
	return t
}
