package main

import (
	"github.com/honkkki/micro-server/logic"
	"github.com/honkkki/micro-server/pb/course"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("api.karina.com.api.course"),
	)

	service.Init()
	err := course.RegisterCourseServiceHandler(service.Server(), &logic.CourseService{})
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(service.Run())
}