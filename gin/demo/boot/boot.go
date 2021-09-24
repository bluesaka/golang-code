package boot

import (
	_ "my-gin/demo/config"
	"my-gin/demo/utils"
	"my-gin/demo/web/controllers"
)

func init() {
	utils.InitZapLog()
	controllers.Init()
}
