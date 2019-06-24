package restapi

import (
	"fmt"
	"net/http"

	"wishbook/restapi/handlers"
	"wishbook/restapi/middleware"

	"github.com/gorilla/mux"
)

// Server properties structure
type Server struct {
	port           string
	guestHandler   handlers.GuestHandler
	userHandler    handlers.UserHandler
	adminHandler   handlers.AdminHandler
	authMiddleware middleware.AuthMiddleware
}

// NewServer is a constructor for server struct
func NewServer(port string, guestHandler handlers.GuestHandler, userHandler handlers.UserHandler, adminHandler handlers.AdminHandler, authMiddleware middleware.AuthMiddleware) *Server {
	return &Server{
		port:           port,
		guestHandler:   guestHandler,
		userHandler:    userHandler,
		adminHandler:   adminHandler,
		authMiddleware: authMiddleware,
	}
}

// RunServer helps to process endpoints
func (s *Server) RunServer() {
	adminRouter := mux.NewRouter()

	adminRouter.HandleFunc("/admin/profile", s.adminHandler.GetAllUsers).Methods(http.MethodGet)
	adminRouter.HandleFunc("/admin/profile", s.adminHandler.AddUser).Methods(http.MethodPost)
	adminRouter.HandleFunc("/admin/profile", s.adminHandler.UpdateUser).Methods(http.MethodPut)

	adminRouter.HandleFunc("/admin/profile/{id}", s.adminHandler.GetUser).Methods(http.MethodGet)
	adminRouter.HandleFunc("/admin/profile/{id}", s.adminHandler.RemoveUser).Methods(http.MethodDelete)

	adminHandler := s.authMiddleware.CheckRole(adminRouter)
	adminHandler = s.authMiddleware.CheckAuth(adminHandler)
	adminHandler = s.authMiddleware.CheckCookie(adminHandler)
	adminHandler = s.authMiddleware.AddHeaderJSON(adminHandler)

	http.Handle("/admin/", adminHandler)
	userRouter := mux.NewRouter()

	userRouter.HandleFunc("/user", s.userHandler.GetCurrentUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/user", s.userHandler.UpdateCurrentUser).Methods(http.MethodPut)

	userRouter.HandleFunc("/user/wish", s.userHandler.GetAllWishes).Methods(http.MethodGet)
	userRouter.HandleFunc("/user/wish", s.userHandler.AddWish).Methods(http.MethodPost)
	userRouter.HandleFunc("/user/wish", s.userHandler.UpdateWish).Methods(http.MethodPatch)

	userRouter.HandleFunc("/user/wish/{id}", s.userHandler.GetWish).Methods(http.MethodGet)
	userRouter.HandleFunc("/user/wish/{id}", s.userHandler.RemoveWish).Methods(http.MethodDelete)

	userHandler := s.authMiddleware.CheckAuth(userRouter)
	userHandler = s.authMiddleware.CheckCookie(userHandler)
	userHandler = s.authMiddleware.AddHeaderJSON(userHandler)
	http.Handle("/user/", userHandler)

	guestRouter := mux.NewRouter()
	guestRouter.HandleFunc("/login", s.guestHandler.Login).Methods(http.MethodPost)
	guestRouter.HandleFunc("/logout", s.guestHandler.Logout)
	http.Handle("/", guestRouter)

	fmt.Println("starting server at", s.port)
	http.ListenAndServe(s.port, nil)
}
