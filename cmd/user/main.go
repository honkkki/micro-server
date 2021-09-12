package main

import (
	"github.com/honkkki/micro-server/model"
	"github.com/honkkki/micro-server/pb/user"
	"github.com/honkkki/micro-server/rpcservice"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"log"
)

func main() {
	var err error
	err = model.InitDB("micro_user")
	if err != nil {
		logger.Error("[init database] failed: ", err)
		return
	}

	service := micro.NewService(
		micro.Name("karina.com.api.user"),
	)

	service.Init()
	err = user.RegisterUserServiceHandler(service.Server(), rpcservice.NewUserService(service.Client()))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(service.Run())
}
