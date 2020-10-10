package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

/*
链接数据库
*/
func Connect(){
	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")         //数据可用户名
	dbPassword := config.String("db_password") //密码
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	//fmt.Println(dbDriver,dbUser,dbPassword)


	//链接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIP+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil { //err不为nil,表示连接数据库时出现了错误，程序就在此中断，不用再执行了
		//早发现，早解决
		panic("数据可链接错误，请检查配置")
	}
	Db = db
}

