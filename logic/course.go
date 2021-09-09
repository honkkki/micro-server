package logic

import (
	"context"
	"github.com/honkkki/micro-server/pb/course"
)

type CourseService struct {
}

func (c *CourseService) GetTop(ctx context.Context, req *course.CourseRequest, resp *course.CourseResponse) error {
	res := make([]*course.Course, 0)
	res = append(res, course.NewCourse(1, "golang"))
	res = append(res, course.NewCourse(2, "grpc"))
	resp.Result = res
	return nil
}