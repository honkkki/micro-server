package model

import (
	"fmt"
	"github.com/honkkki/micro-server/config"
	"github.com/honkkki/micro-server/tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

// user:password@/dbname?charset=utf8&parseTime=True&loc=Local

var Db *gorm.DB

// Model base model
type Model struct {
	ID        int          `gorm:"primary_key" json:"id"`
	CreatedAt tools.MyTime `json:"created_at"`
	UpdatedAt tools.MyTime `json:"updated_at"`
}

func InitDB(dbName string) error {
	var (
		err                        error
		dbType, user, password, ip string
		port, maxIdle, maxOpen     int
	)

	// 从配置中心获取配置
	config.InitConfig()
	dbConfig := config.ConfigData.Db
	dbType = dbConfig.Type
	user = dbConfig.User
	password = dbConfig.Password
	ip = dbConfig.Ip
	port = dbConfig.Port
	maxIdle = dbConfig.MaxIdle
	maxOpen = dbConfig.MaxOpen

	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		ip,
		strconv.Itoa(port),
		dbName))

	if err != nil {
		return err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(maxIdle)
	Db.DB().SetMaxOpenConns(maxOpen)
	return nil
}
