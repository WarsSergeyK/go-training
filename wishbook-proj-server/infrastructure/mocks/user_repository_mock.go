package mocks

import (
	"context"
	"fmt"

	"wishbook/domain"

	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock is mock structure for Unit tests
type UserRepositoryMock struct {
	mock.Mock
}

// InsertUser inserts user
func (mock *UserRepositoryMock) InsertUser(ctx context.Context, user domain.User) (id string, err error) {
	return id, nil
}

// GetUserByID getting user structure by ID
func (mock *UserRepositoryMock) GetUserByID(ctx context.Context, id string) (userResult domain.User, err error) {
	args := mock.Called(id)
	fmt.Println(args)
	return args.Get(0).(domain.User), args.Error(1)
}

// GetAllUsers getting the entire users collection
func (mock *UserRepositoryMock) GetAllUsers(ctx context.Context) (usersResults []*domain.User, err error) {
	return usersResults, nil
}

// GetUserByLogin getting user structure by Login
func (mock *UserRepositoryMock) GetUserByLogin(ctx context.Context, login string) (userResult domain.User, err error) {
	args := mock.Called(login)
	fmt.Println(args)
	return args.Get(0).(domain.User), args.Error(1)
}

// RemoveUserByID removes user by his ID
func (mock *UserRepositoryMock) RemoveUserByID(ctx context.Context, id string) (deleted int64, err error) {
	return deleted, nil
}

// UpdateUser rewrites user by new data
func (mock *UserRepositoryMock) UpdateUser(ctx context.Context, user domain.User) (updated int64, err error) {
	return 0, nil
}
