package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//展示默认的首页；即用户注册页面
func (c *MainController) Get() {
	c.TplName = "register.html"
}
