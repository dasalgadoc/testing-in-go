package api

import (
	"github.com/gin-gonic/gin"
)

func Endpoints(application *Application, router *gin.Engine) {
	router.GET("/ping", application.Controllers.Ping.Pong)
	router.GET("/student/:student_id", application.Controllers.GetStudent.Get)
}
