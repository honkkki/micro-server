package main

import (
	"github.com/gin-gonic/gin"
	"github.com/honkkki/micro-server/httpserver"
	"github.com/honkkki/micro-server/logic/course"
	"github.com/micro/go-micro/v2/web"
	"log"
)

func main() {
	// 引入gin使用http服务
	e := httpserver.InitServer()
	e.GET("/course", func(ctx *gin.Context) {
		course.GetTopList(ctx)
	})

	service := web.NewService(
		web.Name("api.karina.com.http.course"),
		web.Handler(e),
	)

	err := service.Init()
	if err != nil {
		log.Println(err)
		return
	}
	log.Fatal(service.Run())
}
