package main

import (
	"github.com/honkkki/micro-server/pb/user"
	"github.com/honkkki/micro-server/rpcservice"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("api.karina.com.api.user"),
	)

	service.Init()
	err := user.RegisterUserServiceHandler(service.Server(), rpcservice.NewUserService(service.Client()))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(service.Run())
}
