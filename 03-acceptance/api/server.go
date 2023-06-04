package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StartGinServer(application *Application) *http.Server {
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
