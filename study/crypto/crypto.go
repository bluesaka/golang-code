package main

import (
	"fmt"
	"go-code/study/crypto/aes"
)

func main() {
	//fmt.Println(md5.MD5("abc"))
	e, _ := aes.AesEncryptCBC("abc")
	d, _ := aes.AesDecryptCBC(e)
	fmt.Println(e, d)
}


