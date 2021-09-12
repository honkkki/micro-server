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

	logger.Info(model.Db)
	var count int
	model.Db.Table("user").Select("id").Count(&count)
	logger.Info("users count: ", count)

	// 引入gin使用http服务
	e := httpserver.InitServer()
	e.GET("/course", func(ctx *gin.Context) {
		course.GetTopList(ctx)
	})

	service := web.NewService(
		web.Name("karina.com.http.course"),
		web.Handler(e),
	)

	err = service.Init()
	if err != nil {
		log.Println(err)
		return
	}
	log.Fatal(service.Run())
}
