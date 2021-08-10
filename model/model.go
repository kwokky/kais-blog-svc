package model

import (
	"fmt"
	"github.com/kwokky/kais-blog-svc/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

const (
	_maxOpenConn     = 100
	_connMaxLifetime = 10 * time.Second
)

func initDB() {
	conf := config.Get().Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("连接数据库失败", err, conf)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库实例失败", err)
	}

	sqlDB.SetMaxOpenConns(_maxOpenConn)
	sqlDB.SetConnMaxLifetime(_connMaxLifetime)

	autoMigrate()
}

func autoMigrate() {
	_ = db.AutoMigrate(
		&Category{},
		&Post{},
		&PostTag{},
		&Tag{},
	)
}

func DB() *gorm.DB {
	if db == nil {
		initDB()
	}

	return db
}
