package core

import (
	"context"
	"time"

	"wishbook/domain"
	"wishbook/infrastructure"
)

const (
	contextTimeout = 30000
)

// AuthGuard is the interface that keeps user's access checking
type AuthGuard interface {
	IsAuthorized(login string, role domain.UserRole) bool
	IsLoggedIn(login string) bool
	IsPasswordCorrect(login string, passwordHash string) bool
	IsLoginAvailable(login string) bool
}

type authGuard struct {
	usersRepository infrastructure.UserRepository
}

// NewAuthGuard is constructor for authGuard
func NewAuthGuard(usersRepository infrastructure.UserRepository) AuthGuard {
	return &authGuard{
		usersRepository: usersRepository,
	}
}

func (g *authGuard) IsAuthorized(userID string, role domain.UserRole) bool {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Millisecond)
	defer cancel()

	user, err := g.usersRepository.GetUserByID(ctx, userID)
	if err != nil {
		return false
	}
	if user.ID != "" {
		if user.Role == role {
			return true
		}
	}
	return false
}

func (g *authGuard) IsLoggedIn(userID string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Millisecond)
	defer cancel()

	user, err := g.usersRepository.GetUserByID(ctx, userID)
	if err != nil {
		return false
	}
	if user.ID != "" {
		return true
	}
	return false
}

func (g *authGuard) IsPasswordCorrect(login string, passwordHash string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Millisecond)
	defer cancel()

	user, err := g.usersRepository.GetUserByLogin(ctx, login)
	if err != nil {
		return false
	}
	if user.PasswordHash == passwordHash {
		return true
	}
	return false
}

func (g *authGuard) IsLoginAvailable(login string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout*time.Millisecond)
	defer cancel()

	if len(login) == 0 {
		return false
	}

	user, err := g.usersRepository.GetUserByLogin(ctx, login)
	if err != nil {
		return true
	}

	return user.ID == ""
}
