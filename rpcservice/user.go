package rpcservice

import (
	"context"
	"strconv"

	"github.com/honkkki/micro-server/pb/user"
	"github.com/micro/go-micro/v2/client"
)

type UserService struct {
	client client.Client
}

func NewUserService(c client.Client) *UserService {
	return &UserService{client: c}
}

func (u *UserService) Test(ctx context.Context, req *user.UserRequest, resp *user.UserResponse) error {
	resp.Ret = strconv.Itoa(int(req.Id)) + "---karina"
	//s := course.NewCourseService("karina.com.service.course", u.client)
	//res, err := s.GetTop(ctx, &course.CourseRequest{Size: 10})
	//if err != nil {
	//	log.Println("call course service:GetTop failed", err)
	//	return err
	//}
	//log.Println(res.Result)
	return nil
}
