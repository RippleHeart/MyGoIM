package DB

import (
	"fmt"
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
	sqlDB, err := MySQL.DB() // ← 注意这里是 MySQl（你的拼写）
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MySQL连接成功")
	sqlDB.SetMaxIdleConns(Conf.Conf.M.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Conf.Conf.M.MaxOpenConns)
	err = MySQL.AutoMigrate(
		&User{},      // 先建主表
		&UserInfo{},  // 一对一依赖 User
		&Group{},     // 主表
		&UserUser{},  // 中间表依赖 User
		&GroupUser{}, // 中间表依赖 User + Group
	)
	if err != nil {
		log.Fatal("建表失败:", err)
	}
	fmt.Println("表创建成功")

}
