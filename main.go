package main

import (
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {

	/*block0 := blockchain.CerateGenesisBlock()
	block1 := blockchain.NewBlock(block0.Height + 1,block0.Hash,[]byte("a"))
	fmt.Println(block0,block1)
	return*/

	//连接数据库
	db_mysql.Connect()
    //静态资源文件映射（路径）。
    beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

