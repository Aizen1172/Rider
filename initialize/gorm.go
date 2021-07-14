package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"riderStyemProject/global"
	"riderStyemProject/model"
)

var err error

func Gorm() {
	//连接数据库
	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp/rider-system?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171, //默认字符串默认长度  utf8mb4:171   utf8:256
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "r_", //添加表名前缀
			SingularTable: true, //使用单数表名，例如平时默认表名:“riders”，使用后“rider”
		},
		DisableForeignKeyConstraintWhenMigrating: true, //关闭自动创建外键约束(逻辑外键   在代码里体现外键关系)
	})
	if err != nil {
		log.Println("数据库连接失败。")
		os.Exit(1)
	}
	//创建表
	GM := global.DB.Migrator()
	err = GM.AutoMigrate(
		model.User{},
		model.Other{},
		model.Rider{},
		model.File{},
		)
	if err != nil {
		log.Println(err,"创建表失败。")
		os.Exit(1)
	}
}