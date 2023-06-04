package server

import (
	"dasalgadoc.com/go-testing/03-acceptance/api"
	"github.com/gin-gonic/gin"
)

func Endpoints(application *api.Application, router *gin.Engine) {
	router.GET("/ping", application.Controllers.Ping.Pong)
	router.GET("/student/:student_id", application.Controllers.GetStudent.Get)
}
