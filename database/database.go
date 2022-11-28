package database

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"unit-test/log"
)

var Db *gorm.DB

var MysqlUrl = "root:123456@tcp(localhost:3306)/unit_test?charset=utf8mb4&parseTime=True&loc=Asia%2fShanghai&timeout=10s"
var MysqlMaxIdleCount = 10
var MysqlMaxOpenCount = 20

func init() {
	initDb()
}

/**
初始化数据库
*/
func initDb() {
	log.Logger.Info("initDb, start")
	var err error
	Db, err = gorm.Open(mysql.Open(MysqlUrl), &gorm.Config{})
	//无法连接数据库则退出程序
	if err != nil {
		log.Logger.Fatal("gorm.Open error", zap.Error(err))
	}

	//配置数据库连接池
	sqlDb, err := Db.DB()
	if err != nil {
		log.Logger.Fatal("Db.DB error", zap.Error(err))
	}
	sqlDb.SetMaxIdleConns(MysqlMaxIdleCount)
	sqlDb.SetMaxOpenConns(MysqlMaxOpenCount)

	log.Logger.Info("initDb, end")
}
