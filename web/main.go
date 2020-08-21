package main

import (
	"audition/router"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro/web"
)

func main() {
	router := router.InitRouter()
	service := web.NewService(
		web.Name(beego.AppConfig.DefaultString("appname","com.glc.audition")),
		//web.RegisterTTL(time.Second*time.Duration(beego.AppConfig.DefaultInt("registerttl",30))),
		//web.RegisterInterval(time.Second*time.Duration(beego.AppConfig.DefaultInt("registerinterval",20))),
		web.Address(beego.AppConfig.DefaultString("address",":18001")),
		web.Handler(router),
	)
	service.Run()
}
