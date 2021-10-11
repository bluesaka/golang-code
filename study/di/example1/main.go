/**
此例子中没有依赖注入，app依赖了全局变量db。
db对象游离在全局作用于，暴露给包下的其他模块，有被修改的风险
*/

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type App struct{}

// GetData 从MySQL获得数据
func (a *App) GetData(id int) {
	//db.Find("blah blah...")
	fmt.Println("GetData from MySQL")
}

func main() {
	app := new(App)
	app.GetData(1)
}
