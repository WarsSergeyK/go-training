package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"wishbook/domain"
	"wishbook/infrastructure"
)

// GuestHandler has all guest functions
type GuestHandler interface {
	Logout(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type guestHandler struct {
	userRepository infrastructure.UserRepository
}

// NewGuestHandler is constructor for guestHandler
func NewGuestHandler(userRepository infrastructure.UserRepository) GuestHandler {
	return &guestHandler{
		userRepository: userRepository,
	}
}

func (h *guestHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging")

	var u domain.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
		return
	}

	if u.PasswordHash == "" {
		fmt.Println("Empty password")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()

	fu, err := ur.GetUserByLogin(ctx, u.Login)
	if err != nil {
		log.Print(err)
	}

	if strings.Compare(u.PasswordHash, fu.PasswordHash) != 0 {
		fmt.Println("Incorrect password")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("Correct password")

	expiration := time.Now().Add(cookieExpirationH * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   fu.ID,
		Expires: expiration,
	}

	http.SetCookie(w, &cookie)
	//http.Redirect(w, r, "/user/wish", http.StatusFound)
}

func (h *guestHandler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In logout")
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return
	}

	fmt.Println(session.Expires)
	session.Expires = time.Now().AddDate(0, 0, -1)

	http.SetCookie(w, session)
}
