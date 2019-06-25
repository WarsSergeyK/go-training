package infrastructure

import (
	"context"
	"log"
	"time"

	"wishbook/domain"
)

// Check for admin account existance. Create if needed
func init() {

	adminLogin := "admin"

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Millisecond)
	defer cancel()

	log.Println("Check admin db account")

	g := NewUserRepository()

	user, err := g.GetUserByLogin(ctx, adminLogin)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	if user.ID == "" {
		u := domain.User{
			Login:        adminLogin,
			PasswordHash: "admin",
			Role:         domain.Admin,
		}

		_, err := g.InsertUser(ctx, u)
		if err != nil {
			log.Printf("error: %v\n", err)
			return
		}
	}
}
