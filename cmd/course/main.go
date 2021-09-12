package main

import (
	"github.com/honkkki/micro-server/pb/course"
	"github.com/honkkki/micro-server/rpcservice"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("karina.com.api.course"),
	)

	service.Init()
	err := course.RegisterCourseServiceHandler(service.Server(), rpcservice.NewCourseService(service.Client()))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(service.Run())
}
