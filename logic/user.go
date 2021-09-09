package logic

import (
	"context"
	"github.com/honkkki/micro-server/pb/user"
	"strconv"
)

type UserService struct {
}

func (u *UserService) Test(ctx context.Context, req *user.UserRequest, resp *user.UserResponse) error {
	resp.Ret = strconv.Itoa(int(req.Id)) + "---karina"
	return nil
}