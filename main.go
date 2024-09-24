package main

import (
	"github.com/ananascharles/binify/routes"
)

func main() {
	router := routes.SetupRouter()

	if err := router.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}

}
