package main

import myjwt "go-code/study/jwt"

func main() {
	token, _ := myjwt.CreateJwt(myjwt.User{Account: "test3", Age: 22})
	_, _ = myjwt.ParseJwt(token)

}
