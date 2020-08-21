package handler

import (
	"audition/models"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func GetMultiple(context *gin.Context){
	count := models.GetJavaAuditionCount()
	logs.Debug(count)
	rand.Seed(time.Now().UnixNano())
	var ids []int64
	if count > 10 {
		for {
			if len(ids) < 10 {

				id := rand.Intn(count)+523
				if !isContain(ids, int64(id)) && id != 0 {
					ids = append(ids, int64(id))
				}
			} else {
				break
			}
		}
	}
	audition := models.GegTenAudition(ids)
	for _, v := range audition {
		logs.Info("%+v", v)
	}
	context.JSON(200,&audition)
}

func isContain(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
