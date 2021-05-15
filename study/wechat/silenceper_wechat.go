/**
@link https://github.com/silenceper/wechat

*/

package wechat

import (
	"github.com/silenceper/wechat/v2"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"go-code/study/redis/redigo"

)

func Wechat1() {
	wc := wechat.NewWechat()
	redigo.GetRedis()
	cfg := &offConfig.Config{
		AppID: "",
		AppSecret: "",
		Token:     "xxx",
		//EncodingAESKey: "xxxx",
		//Cache: memory,
	}
	_ = wc.GetOfficialAccount(cfg)

}
