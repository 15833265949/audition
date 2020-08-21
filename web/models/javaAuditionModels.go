package models

import (
	"audition/models/mysql"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type Radio struct {
	gorm.Model
	Title       string `gorm:"size:255"`
	Type        int
	OptionA     string `gorm:"size:255"`
	OptionB     string `gorm:"size:255"`
	OptionC     string `gorm:"size:255"`
	OptionD     string `gorm:"size:255"`
	Answer      string `gorm:"size:10"`
	Explanation string `gorm:"size:255"`
}

func GetJavaAuditionCount() (count int) {
	db, err := mysql.GetConnect()
	defer db.Close()
	if err != nil {
		logs.Debug("err:%s 数据库连接错误", err)
	}
	db.Table("radio").Count(&count)
	return
}

func GegTenAudition(ids []int64) (ras []*Radio) {
	db, err := mysql.GetConnect()
	defer db.Close()
	if err != nil {
		logs.Debug("err:%s 数据库连接错误", err)
	}
	var radios []*Radio

	db.Debug().Where(ids).Find(&radios)


	return radios
}

