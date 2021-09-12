package course

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/honkkki/micro-server/pb/course"
	"github.com/micro/go-micro/v2/client/grpc"
)

func GetTopList(ctx *gin.Context) {
	c := grpc.NewClient()
	// 调用grpc服务
	s := course.NewCourseService("karina.com.api.course", c)
	res, err := s.GetTop(context.Background(), &course.CourseRequest{Size: 10})
	if err != nil {
		ctx.JSON(500, gin.H{
			"data": nil,
		})
	} else {
		ctx.JSON(200, gin.H{
			"data": res.Result,
		})
	}
}
