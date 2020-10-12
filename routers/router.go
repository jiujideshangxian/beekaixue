package routers

import (
	"DataCertPaltPhone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册
    beego.Router("/register",&controllers.ResgiterController{})
    beego.Router("/login",&controllers.LoginController{})
}
