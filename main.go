package main

import (
	"dasalgadoc.com/go-testing/03-acceptance/api"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Starting API...")

	application := api.BuildApplication()
	server := api.StartGinServer(application)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
