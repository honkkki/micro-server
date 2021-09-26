package main

import (
	"github.com/gin-gonic/gin"
	"github.com/honkkki/micro-server/httpserver"
	"github.com/honkkki/micro-server/logic/course"
	"github.com/honkkki/micro-server/model"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"log"
)

func main() {
	var err error
	// init db
	err = model.InitDB("micro_user")
	if err != nil {
		logger.Error("[init database] failed: ", err)
		return
	}

	var user model.User
	model.Db.Table("user").Select("*").First(&user)
	log.Println(user.Nickname)

	// 引入gin使用http服务
	e := httpserver.InitServer()
	e.GET("/courselist", func(ctx *gin.Context) {
		course.GetTopList(ctx)
	})

	service := web.NewService(
		web.Name("karina.com.api.coursehttp"),
		web.Version("1.0.0"),
		web.Handler(e),
	)

	err = service.Init()
	if err != nil {
		log.Println(err)
		return
	}
	log.Fatal(service.Run())
}
