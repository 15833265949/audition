package models

import (
	"audition/models/mysql"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:30"`
	NickName string `gorm:"size:20"`
	Phone    string `gorm:"size:20"`
}

func GetBoolByName(name string) bool {
	db, err := mysql.GetConnect()
	if err != nil {
		logs.Error(err)
	}
	count := 0
	db.Debug().Model(&User{}).Where("nick_name = ?",name).Count(&count)
	if count > 0{
		return false
	}else {
		return true
	}
}

func GetUserByName(name string) *User {
	db, err := mysql.GetConnect()
	if err != nil {
		logs.Error(err)
	}
	user := &User{}
	db.Debug().Where("nick_name = ?",name).First(user)
	return user
}

func InsertUser(name ,password,email,phone string){
	db, err := mysql.GetConnect()
	if err != nil {
		logs.Error(err)
	}
	user := User{
		NickName: name,
		Password: password,
		Email: email,
		Phone: phone,
	}

	db.Debug().Create(&user)
}
