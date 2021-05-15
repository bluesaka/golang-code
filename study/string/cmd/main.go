package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"hash"
	"sort"
	"strings"
)

func main() {
	sortString1()
}

func string1() {
	s := "type=1&server_name=xxx&account=xxx&nickname=xxx&store_type=1&money=1&coin=1"
	ss := strings.Split(s, "&")
	fmt.Println(ss)

	sss := make(map[string]string, len(ss))

	for _, v := range ss {
		vv := strings.Split(v, "=")
		sss[vv[0]] = vv[1]
	}
	fmt.Println(sss)

	b, _ := jsoniter.Marshal(sss)
	fmt.Println(string(b))
}

func sortString1() {
	//sig := "sig"
	token := "token"
	timestamp := "123"
	nonce := "nonce"

	s := []string{token, timestamp, nonce}
	sort.Strings(s)
	fmt.Println(s)

	ss := strings.Join(s, "")
	fmt.Println(ss)

	h := sha1.New()
	h.Write([]byte(ss))
	e := hex.EncodeToString(h.Sum(nil))
	fmt.Println(e)

	e2 := GenerateSHA1("sha1",  ss)
	fmt.Println(e2)
}

func GenerateSHA1(way string, data string) string {

	var h hash.Hash
	if way == "sha256" {
		h = sha256.New()
	} else {
		h = sha1.New()
	}
	// Create a new HMAC by defining the hash type and the key (as byte array)
	// Write Data to it
	h.Write([]byte(data))
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
