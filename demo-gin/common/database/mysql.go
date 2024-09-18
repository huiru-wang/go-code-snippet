package mysql

import (
	"fmt"
	"go-gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var con *gorm.DB
var err error

func Session() *gorm.DB {
	return con
}
func InitMysql() {

	ds := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		ds.Username, //root
		ds.Password, //root
		ds.Host,     //localhost
		ds.Port,     //3306
		ds.Database, //messdb
		ds.Charset,  //utf8mb4
	)
	log.Printf("datasource:%s", dsn)
	con, err = gorm.Open(
		mysql.New(mysql.Config{
			// mysql配置
			DSN: dsn,
		}),
		&gorm.Config{
			// gorm配置
			PrepareStmt: true,           // 创建并缓存预编译语句，可以提高后续的调用速度
			Logger:      NewSqlLogger(), // 自定义sql日志
		})

	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	// 连接池配置
	db, _ := con.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接的最大可复用时间。
	db.SetConnMaxLifetime(10 * time.Second)
}

func NewSqlLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)
}
