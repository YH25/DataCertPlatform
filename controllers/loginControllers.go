package controllers

import (
	"DataCertPlatform/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
//直接跳转展示用户登录页面
func (l *LoginController) Get() {
    l.TplName = "login.html"
}

/*
post方法处理用户的登录请求
 */
func (l *LoginController) Post() {
	//1.解析客户端提交的登录数据
	var User models.User
	err :=l.ParseForm(&User)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录信息解析失败，请重试")
		return
	}
	//2.根据解析到的数据，执行数据库查询操作
	u, err := User.QueryUser()
	//3.判断数据库查询结果
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)//把脱敏的密码的MD5值重新赋值为密码进行

	if err != nil {
		//sql:no rows in result set
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试")
		return
	}

	//4.根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone//动态数据设置
	l.TplName = "home.html"//文件上传界面

}
