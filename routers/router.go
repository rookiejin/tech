package routers

import (
	"tech/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/category" ,&controllers.CategoryController{})
}
