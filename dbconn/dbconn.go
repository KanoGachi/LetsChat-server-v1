package dbconn

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"

	_ "github.com/go-sql-driver/mysql"
)

type AppConf struct {
	DBConfig DBConf `ini:"db"`
}

type DBConf struct {
	Destination string `ini:"mysqldsn"`
}

var (
	Database  *sqlx.DB // 建立连接后可以使用该全局变量操作数据库
	AppConfig AppConf  // 配置信息
)

func InitDB() (err error) {
	cfg, err := ini.Load("config/webapp.ini")
	if err != nil {
		return
	}
	err = cfg.MapTo(&AppConfig)
	if err != nil {
		return
	}
	dsn := AppConfig.DBConfig.Destination
	Database, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	err = Database.Ping()
	if err != nil {
		return
	}
	return
}
