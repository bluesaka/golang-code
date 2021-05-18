package model

type CarInfoReq struct {
	ID   int    // default `p:"id"`
	Area int    `p:"area" v:"required|min:1#area is empty|area要大于0"`
	Pass string `p:"password" v:"required|length:6,10#请输入密码|密码长度为6到10位数"`
}

type CarInfoResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
