package middleware

import (
	"fmt"
	"net/http"

	"wishbook/core"
	"wishbook/domain"
)

// AuthMiddleware interface checking access
type AuthMiddleware interface {
	CheckAuth(next http.Handler) http.Handler
	CheckRole(next http.Handler) http.Handler
	CheckCookie(next http.Handler) http.Handler
	AddHeaderJSON(next http.Handler) http.Handler
}

type authMiddleware struct {
	authGuard core.AuthGuard
}

// NewAuthMiddleware is constructor for authMiddleware structure
func NewAuthMiddleware(authGuard core.AuthGuard) AuthMiddleware {
	return &authMiddleware{
		authGuard: authGuard,
	}
}

func (middleware *authMiddleware) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking auth for:", r.URL.Path)

		session, err := r.Cookie("session_id")
		if err != nil {
			return
		}

		userID := session.Value

		if !middleware.authGuard.IsLoggedIn(userID) {
			fmt.Println("Access denied")
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		fmt.Println("Access granted")

		next.ServeHTTP(w, r)
	})
}

func (middleware *authMiddleware) CheckRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err != nil {
			return
		}

		userID := session.Value

		if !middleware.authGuard.IsAuthorized(userID, domain.Admin) {
			fmt.Println("Access denied")
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		fmt.Println("Access granted")

		next.ServeHTTP(w, r)
	})
}

func (middleware *authMiddleware) CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking for a cookie:", r.URL.Path)

		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			fmt.Println(err)
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}

		if session == nil {
			fmt.Println("No value")
			return
		}
		fmt.Println("Have a cookie")

		w.Header().Set("session_id", session.Value)

		next.ServeHTTP(w, r)
	})
}

func (middleware *authMiddleware) AddHeaderJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Adding JSON header")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
