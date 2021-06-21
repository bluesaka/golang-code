/**
json-iterator 效率比原生的encoding/json 效率高很多，推荐使用
*/
package json

import (
	"github.com/json-iterator/go"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var jsonObj = jsoniter.ConfigCompatibleWithStandardLibrary

func MyMarshal(u User) ([]byte, error) {
	//b, err := jsoniter.Marshal(u)
	b, err := jsonObj.Marshal(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return b, nil
}

func MyUnmarshal(b []byte) User {
	var u User
	if err := jsoniter.Unmarshal(b, &u); err != nil {
		log.Println(err)
	}
	return u
}
