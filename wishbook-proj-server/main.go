package main

import (
	"wishbook/core"
	"wishbook/infrastructure"
	"wishbook/restapi"
	"wishbook/restapi/handlers"
	"wishbook/restapi/middleware"
)

const (
	httpServerPort = ":8080"
)

func main() {

	userRepository := infrastructure.NewUserRepository()

	authGuard := core.NewAuthGuard(userRepository)
	authMiddleware := middleware.NewAuthMiddleware(authGuard)

	guestHandler := handlers.NewGuestHandler(userRepository)
	userHandler := handlers.NewUserHandler(userRepository)
	adminHandler := handlers.NewAdminHandler(userRepository)

	server := restapi.NewServer(httpServerPort, guestHandler, userHandler, adminHandler, authMiddleware)
	server.RunServer()
}
