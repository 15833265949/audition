package models

import (
	"audition/models/mysql"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"size:30"`
	Account  string `gorm:"size:30"`
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
	db.Debug().Model(&User{}).Where("account = ?", name).Count(&count)
	if count > 0 {
		return false
	} else {
		return true
	}
}

func GetUserByName(account string) *User {
	db, err := mysql.GetConnect()
	if err != nil {
		logs.Error(err)
	}
	user := &User{}
	db.Debug().Where("account = ?", account).First(user)
	return user
}

func InsertUser(name,account, password, email, nickname, phone string) {
	db, err := mysql.GetConnect()
	if err != nil {
		logs.Error(err)
	}
	user := User{
		UserName: name,
		Account: account,
		Password: password,
		Email:    email,
		NickName: nickname,
		Phone:    phone,
	}

	db.Debug().Create(&user)
}
