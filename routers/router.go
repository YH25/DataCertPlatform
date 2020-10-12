package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//ruoter;路由
    beego.Router("/", &controllers.MainController{})
    //用户注册接口
    beego.Router("/register",&controllers.RegisterController{})
    //用户登录接口
    beego.Router("/enter",&controllers.EnterControllers{})
}
