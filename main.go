package main

import (
	"fmt"
	"server-graphql/internal/services"
	"server-graphql/internal/util"
)

func main() {
	fmt.Println("Hello, I'm H.")

	service, err := services.InitGraphQLServices()
	if err != nil {
		util.HandleFatalf(err)
	}
	service.ListenAndServe()
}
