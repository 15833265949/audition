package handler

import (
	"audition/models"
	"audition/utils"
	"github.com/gin-gonic/gin"
)

func Registered(c *gin.Context)  {
	nickName:= c.Query("nickname")
	pwd := c.Query("password")
	userName := c.Query("username")
	account := c.Query("account")

	if pwd == ""{
		utils.WriteRsp(c,1,"密码为空，请输入密码！",nil)
		return
	}
	if account == ""{
		utils.WriteRsp(c,1,"账号为空，请输入账号！",nil)
		return
	}

	email := c.PostForm("email")
	phone := c.PostForm("phone")

	b := models.GetBoolByName(account)
	if !b {
		utils.WriteRsp(c,1,"账号已存在！",nil)
		return
	}

	models.InsertUser(userName,account,pwd,email,nickName,phone)
	utils.WriteRsp(c,0,"注册成功！",nil)

}

func Login(c *gin.Context){
	account:= c.Query("account")
	pwd := c.Query("password")

	if pwd == ""{
		utils.WriteRsp(c,1,"密码为空，请输入密码！",nil)
		return
	}
	if account == ""{
		utils.WriteRsp(c,1,"账号为空，请输入账号！",nil)
		return
	}

	user := models.GetUserByName(account)
	if user != nil && user.Account != "" && user.Password != ""{
		if pwd == user.Password{
			utils.WriteRsp(c,0,"登录成功",user)
		}else {
			utils.WriteRsp(c,1,"密码错误",nil)
		}
	}else {
		utils.WriteRsp(c,1,"用户名不存在",nil)
	}

}


