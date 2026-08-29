package DB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mygoim/Conf"
)

var MySQL *gorm.DB

func InitMySQL() {
	var err error
	MySQL, err = gorm.Open(mysql.Open(Conf.DSN), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("MySQL连接失败:", err)
	}
	sqlDB, err := MySQL.DB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MySQL连接成功")
	sqlDB.SetMaxIdleConns(Conf.Conf.M.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Conf.Conf.M.MaxOpenConns)
	//err = MySQL.AutoMigrate(
	//	&User{},
	//	&UserInfo{},
	//	&Group{},
	//	&UserUser{},
	//	&GroupUser{},
	//)
	//if err != nil {
	//log.Fatal("建表失败:", err)
	//}

}
