package DBInit

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mygoim/Conf"
)

func InitMySQl() {

	db, err := gorm.Open(mysql.Open(Conf.DSN), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("MySQL连接失败:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(Conf.Conf.M.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Conf.Conf.M.MaxOpenConns)

	log.Println("MySQL连接成功")
}
