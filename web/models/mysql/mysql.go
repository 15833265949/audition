package mysql

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	dbHandler *gorm.DB
	addr      string
)

func GetConnString() string {
	if addr == "" {
		ip := "127.0.0.1"
		port := "3306"
		username := "root"
		passwd := "111"
		dbname := "audition"
		addr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local&parseTime=true", username, passwd, ip, port, dbname)
	}
	return addr
}

func GetConnect() (*gorm.DB, error) {
	if dbHandler != nil {
		err := dbHandler.DB().Ping()
		if err == nil {
			return dbHandler, nil
		}
	}
	var err error
	dbHandler, err = gorm.Open("mysql", GetConnString())
	if err != nil {
		logs.Error("err:%s err:%s", err, addr)
		return nil, err
	}
	// 全局禁用表名复数
	dbHandler.SingularTable(true)
	// SetMaxIdleConns 最大空闲连接数据
	dbHandler.DB().SetMaxIdleConns(beego.AppConfig.DefaultInt("mysqlMaxidel", 4))

	// SetMaxOpenConns 最大使用连接数
	dbHandler.DB().SetMaxOpenConns(beego.AppConfig.DefaultInt("mysqlMaxopen", 5))

	// SetConnMaxLifetime 设置连接保持的最大时间
	defaultInt := beego.AppConfig.DefaultInt("mysqlTimeout", 10)
	dbHandler.DB().SetConnMaxLifetime(time.Duration(defaultInt) * time.Second)
	return dbHandler, nil
}
