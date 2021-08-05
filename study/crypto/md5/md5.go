package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	//fmt.Printf("%x", h.Sum(nil))
	//return fmt.Sprintf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5_2(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	//fmt.Printf("%x", h.Sum(nil))
	//return fmt.Sprintf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5_3(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}