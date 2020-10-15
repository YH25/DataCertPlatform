package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//展示默认的首页；即用户注册页面
func (c *MainController) Get() {
	fmt.Println("helloword")
	c.TplName = "register.html"
}
