package handler

import (
	"audition/models"
	"audition/utils"
	"github.com/gin-gonic/gin"
)

func Registered(c *gin.Context)  {
	rsp := &utils.Rsp{}
	nickName:= c.Query("name")
	pwd := c.Query("password")

	if pwd == ""{
		utils.WriteRsp(c,1,"密码为空，请输入密码！",nil)
		return
	}
	if nickName == ""{
		utils.WriteRsp(c,1,"用户名为空，请输入用户名！",nil)
		return
	}

	email := c.PostForm("email")
	phone := c.PostForm("phone")

	b := models.GetBoolByName(nickName)
	if !b {
		rsp.Code=1
		rsp.Msg = "用户名已存在！"
		c.JSON(200,gin.H{
			"data":rsp,
		})
		return
	}

	models.InsertUser(nickName,pwd,email,phone)
	utils.WriteRsp(c,0,"注册成功！",nil)

}

func Login(c *gin.Context){
	nickName:= c.Query("name")
	pwd := c.Query("password")

	if pwd == ""{
		utils.WriteRsp(c,1,"密码为空，请输入密码！",nil)
		return
	}
	if nickName == ""{
		utils.WriteRsp(c,1,"用户名为空，请输入用户名！",nil)
		return
	}

	user := models.GetUserByName(nickName)
	if user != nil && user.NickName != "" && user.Password != ""{
		if pwd == user.Password{
			utils.WriteRsp(c,0,"登录成功",nil)
		}else {
			utils.WriteRsp(c,1,"密码错误",nil)
		}
	}else {
		utils.WriteRsp(c,1,"用户名不存在",nil)
	}

}


