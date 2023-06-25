package Core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"web/Global"
)

func InitGorm() *gorm.DB {
	dsn := Global.Config.Mysql.Dsn()
	var mylog logger.Interface
	if Global.Config.System.Env == "debug" {
		mylog = logger.Default.LogMode(logger.Info)
	} else {
		mylog = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mylog,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s]mysql连接失败", dsn))
	}
	//fmt.Println("mysql连接成功")
	return db
}
