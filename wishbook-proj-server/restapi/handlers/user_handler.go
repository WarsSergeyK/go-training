package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"wishbook/domain"
	"wishbook/infrastructure"

	"github.com/gorilla/mux"
)

// UserHandler has all users functions
type UserHandler interface {
	GetCurrentUser(w http.ResponseWriter, r *http.Request)
	UpdateCurrentUser(w http.ResponseWriter, r *http.Request)

	GetAllWishes(w http.ResponseWriter, r *http.Request)
	GetWish(w http.ResponseWriter, r *http.Request)
	AddWish(w http.ResponseWriter, r *http.Request)
	UpdateWish(w http.ResponseWriter, r *http.Request)
	RemoveWish(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userRepository infrastructure.UserRepository
}

// NewUserHandler is constructor
func NewUserHandler(userRepository infrastructure.UserRepository) UserHandler {
	return &userHandler{
		userRepository: userRepository,
	}
}

func (h *userHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println("error:", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, string(JSON))
	} else {
		fmt.Printf("Nothing found\n")
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

func (h *userHandler) UpdateCurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewUserRepository()
	var u domain.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
		return
	}

	updated, err := ur.UpdateUser(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents\n", updated)
}

func (h *userHandler) GetAllWishes(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewWishRepository()

	allWishes, err := ur.GetAllActualWishes(ctx)
	if err != nil {
		log.Fatal(err)
	}

	JSON, err := json.MarshalIndent(allWishes, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSON))
}

func (h *userHandler) GetWish(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewWishRepository()

	fu, err := ur.GetWishByID(ctx, id)
	if err != nil {
		log.Print(err)
	}

	if fu.ID != "" {
		fmt.Printf("Found a single document: %+v\n", fu)

		JSON, err := json.MarshalIndent(fu, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w, string(JSON))
	} else {
		fmt.Printf("Nothing found\n")
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

func (h *userHandler) AddWish(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewWishRepository()
	var u domain.Wish

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := ur.InsertWish(ctx, u)
	if err != nil {
		log.Fatal(err)
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

func (h *userHandler) UpdateWish(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewWishRepository()
	var u domain.Wish

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
		return
	}

	updated, err := ur.UpdateWish(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents\n", updated)
	js, err := json.Marshal(struct {
		Val int64 `json:"Updated"`
	}{updated})
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write(js)
}

func (h *userHandler) RemoveWish(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeoutMs*time.Millisecond)
	defer cancel()

	ur := infrastructure.NewWishRepository()

	deleted, err := ur.RemoveWishByID(ctx, id)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "406 not acceptable", http.StatusNotAcceptable)
		return
	}

	fmt.Printf("Deleted %v documents in the wishes collection\n", deleted)

	fmt.Printf("Updated %v documents\n", deleted)
	js, err := json.Marshal(struct {
		Val int64 `json:"Deleted"`
	}{deleted})
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write(js)
}
