package server

import (
	"dasalgadoc.com/go-testing/03-acceptance/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StartGinServer(application *api.Application) *http.Server {
	router := gin.New()
	Endpoints(application, router)
	port := application.Config.Api.Port
	if port == "" {
		port = "8081"
	}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
