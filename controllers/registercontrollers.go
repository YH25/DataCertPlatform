package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {

	//1.解析用户端提交的请求数据
	var User models.User
	err :=r.ParseForm(&User)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析错误")
		return
	}
	//2.将解析到的数据保存到数据库中
	 _, err = User.AddUser()
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试")
		return

	}
	//3.将处理结果返回给客户端浏览器
	  //3.1如果成功，跳转登录页面
	   //tpl:template:模板
	   r.TplName = "login.html"
	  //3.2如果失败，提示错误信息

}
