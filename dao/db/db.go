package db

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func Init() error {
	var err error
	//dns := beego.AppConfig.String("MysqlUser") + ":" + beego.AppConfig.String("MysqlPassword") + "@" +
	//	beego.AppConfig.String("MysqlProtocol") + "(" + beego.AppConfig.String("MysqlAddr") + ":" + beego.AppConfig.String("MysqlPort") + ")" + "/" + beego.AppConfig.String("MysqlDatabase") + "?parseTime=true"
	dns := beego.AppConfig.String("MysqlUser") + ":" + ""+ "@" +
		beego.AppConfig.String("MysqlProtocol") + "(" + beego.AppConfig.String("MysqlAddr") + ":" + beego.AppConfig.String("MysqlPort") + ")" + "/" + beego.AppConfig.String("MysqlDatabase") + "?parseTime=true"

	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	err = DB.Ping() //防止数据库账号密码不对没有报错
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}
