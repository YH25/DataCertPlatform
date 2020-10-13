package main

import (
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	//连接数据库
	db_mysql.Connect()
    //静态资源文件映射。
    beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

