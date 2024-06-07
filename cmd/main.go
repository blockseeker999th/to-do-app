package main

import (
	"log"

	todo "github.com/blockseeker999th/to-do-app"
	"github.com/blockseeker999th/to-do-app/internal/handlers"
)

func main() {
	routes := new(handlers.Handler)

	srv := new(todo.Server)

	if err := srv.Run("8080", routes.InitRoutes()); err != nil {
		log.Fatalf("error starting a server on %s", err.Error())
	}

}
