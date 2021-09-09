package logic

import (
	"context"
	"github.com/honkkki/micro-server/pb/course"
	"github.com/honkkki/micro-server/pb/user"
	"github.com/micro/go-micro/v2/client"
	"log"
	"strconv"
)

type UserService struct {
	client client.Client
}

func NewUserService(c client.Client) *UserService {
	return &UserService{client: c}
}

func (u *UserService) Test(ctx context.Context, req *user.UserRequest, resp *user.UserResponse) error {
	resp.Ret = strconv.Itoa(int(req.Id)) + "---karina"
	s := course.NewCourseService("api.karina.com.api.course", u.client)
	res, err := s.GetTop(ctx, &course.CourseRequest{Size: 10})
	if err != nil {
		log.Println("call course service:GetTop failed", err)
	}
	log.Println(res.Result)
	return nil
}