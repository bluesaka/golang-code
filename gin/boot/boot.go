package boot

import (
	_ "my-gin/config"
	"my-gin/utils"
	"my-gin/web/controllers"
)

func init() {
	utils.InitZapLog()
	controllers.Init()
}
