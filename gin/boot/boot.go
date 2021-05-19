package boot

import (
	_ "my-gin/config"
	utils "my-gin/utils/logger"
	"my-gin/web/controllers"
)

func init() {
	utils.InitZapLog()
	controllers.Init()
}
