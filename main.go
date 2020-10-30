package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0 := blockchain.CerateGenesisBlock()//创世区块
	//fmt.Println(block0)
	block1 := blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})
	fmt.Printf("block0的哈希:%x\n",block0.Hash)
	fmt.Printf("block的哈希:%x\n",block1.Hash)
	fmt.Printf("block1de prevHash:%x\n",block1.PervHash)

	//序列化
	blockJson, _ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block:",string(blockJson))

	blockXml, _ := xml.Marshal(block0)
	fmt.Println("通过xml序列化以后的block:", string(blockXml))

	//连接数据库
	db_mysql.Connect()
    //静态资源文件映射（路径）。
    beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

