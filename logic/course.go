package logic

import (
	"context"
	"github.com/honkkki/micro-server/pb/course"
	"github.com/micro/go-micro/v2/client"
)

type CourseService struct {
	client client.Client
}

func NewCourseService(c client.Client) *CourseService {
	return &CourseService{client: c}
}

func (c *CourseService) GetTop(ctx context.Context, req *course.CourseRequest, resp *course.CourseResponse) error {
	res := make([]*course.Course, 0)
	res = append(res, course.NewCourse(1, "golang"))
	res = append(res, course.NewCourse(2, "grpc"))
	resp.Result = res
	return nil
}