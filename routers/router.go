package routers

import (
	"github.com/astaxie/beego"
	"myBeegoPro/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
