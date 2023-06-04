package main

import (
	"dasalgadoc.com/go-testing/03-acceptance/api"
	"dasalgadoc.com/go-testing/03-acceptance/infrastructure/server"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Starting API...")

	application := api.BuildApplication()
	server := server.StartGinServer(application)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
