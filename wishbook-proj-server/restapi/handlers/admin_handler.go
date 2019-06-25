package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"wishbook/core"
	"wishbook/domain"
	"wishbook/infrastructure"

	"github.com/gorilla/mux"
)

// AdminHandler has all admin functions
type AdminHandler interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	RemoveUser(w http.ResponseWriter, r *http.Request)
}

type adminHandler struct {
	userRepository infrastructure.UserRepository
}

// NewAdminHandler is constructor
func NewAdminHandler(userRepository infrastructure.UserRepository) AdminHandler {
	return &adminHandler{
		userRepository: userRepository,
	}
}

func (h *adminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()

	allUsers, err := ur.GetAllUsers(ctx)
	if err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	JSON, err := json.MarshalIndent(allUsers, "", "\t")
	if err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(JSON))
}

func (h *adminHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()

	fu, err := ur.GetUserByID(ctx, id)
	if err != nil {
		log.Print(err)
	}

	if fu.ID != "" {
		fmt.Printf("Found a single document: %+v\n", fu)

		JSON, err := json.MarshalIndent(fu, "", "\t")
		if err != nil {
			log.Printf("error: %v\n", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, string(JSON))
	} else {
		fmt.Printf("Nothing found\n")
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

func (h *adminHandler) AddUser(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()
	var u domain.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	ag := core.NewAuthGuard(ur)
	if ag.IsLoginAvailable(u.Login) == false {
		log.Printf("Not available login\n")
		http.Error(w, "406 not acceptable", http.StatusNotAcceptable)
		return
	}

	id, err := ur.InsertUser(ctx, u)
	if err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Inserted a single document:", id)
	js, err := json.Marshal(struct {
		ID string `json:"ID"`
	}{id})
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write(js)
}

func (h *adminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()
	var u domain.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	updated, err := ur.UpdateUser(ctx, u)
	if err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Updated %v documents\n", updated)

	js, err := json.Marshal(struct {
		Val int64 `json:"Message"`
	}{updated})
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write(js)
}

func (h *adminHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()

	deleted, err := ur.RemoveUserByID(ctx, id)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "406 not acceptable", http.StatusNotAcceptable)
		return
	}
	fmt.Printf("Deleted %v documents in the users collection\n", deleted)

	js, err := json.Marshal(struct {
		Val int64 `json:"Deleted"`
	}{deleted})
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write(js)
}
