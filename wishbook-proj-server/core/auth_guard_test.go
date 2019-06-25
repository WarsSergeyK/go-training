package core

import (
	"errors"
	"testing"

	"wishbook/domain"
	"wishbook/infrastructure"
	"wishbook/infrastructure/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAuthGuard(t *testing.T) {
	ag := NewAuthGuard(infrastructure.NewUserRepository())
	assert.NotNil(t, ag)
}

func TestIsAuthorized_ShouldBeAuthorized(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	// Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByID", mock.Anything).Return(user, nil)

	idForTest := "5d10c26392ab2a29a9df1ab8"

	ag := NewAuthGuard(userRepoMock)

	expectedResult := true
	actualResult := ag.IsAuthorized(idForTest, domain.Admin)

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsAuthorized_Error(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	// Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByID", mock.Anything).Return(user, errors.New("someError"))

	idForTest := "5d10c26392ab2a29a9df1ab8"

	ag := NewAuthGuard(userRepoMock)

	expectedResult := false
	actualResult := ag.IsAuthorized(idForTest, domain.Admin)

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsAuthorized_ShouldNotBeAuthorized(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	// Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "userName",
		Surname:      "userSurname",
		Login:        "user",
		PasswordHash: "useruser",
		Role:         domain.DefaultUser,
	}

	userRepoMock.On("GetUserByID", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	idForTest := "5d10c26392ab2a29a9df1ab8"
	actualResult := ag.IsAuthorized(idForTest, domain.Admin)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}
func TestIsLoggedIn_ShouldBeLoggedIn(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	// Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByID", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	idForTest := "5d10c26392ab2a29a9df1ab8"
	actualResult := ag.IsLoggedIn(idForTest)

	expectedResult := true

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsLoggedIn_ShouldNotBeLoggedIn(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	// Mocked user to get
	user := domain.User{}

	userRepoMock.On("GetUserByID", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	idForTest := "fakeID"
	actualResult := ag.IsLoggedIn(idForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsLoggedIn_Error(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	// Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByID", mock.Anything).Return(user, errors.New("someError"))

	ag := NewAuthGuard(userRepoMock)

	idForTest := "5d10c26392ab2a29a9df1ab8"
	actualResult := ag.IsLoggedIn(idForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsPasswordCorrect_ShouldBeCorrect(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	//	Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByLogin", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	loginForTest := "admin"
	PasswordHashForTest := "adminadmin"
	actualResult := ag.IsPasswordCorrect(loginForTest, PasswordHashForTest)

	expectedResult := true

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsPasswordCorrect_ShouldBeIncorrect(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	//	Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByLogin", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	loginForTest := "admin"
	PasswordHashForTest := "fakePassword"
	actualResult := ag.IsPasswordCorrect(loginForTest, PasswordHashForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsPasswordCorrect_Error(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	//	Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByLogin", mock.Anything).Return(user, errors.New("someError"))

	ag := NewAuthGuard(userRepoMock)

	loginForTest := "admin"
	PasswordHashForTest := "fakePassword"
	actualResult := ag.IsPasswordCorrect(loginForTest, PasswordHashForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsLoginAvailable_ShouldBeAvailable(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	//	Mocked user to get
	user := domain.User{}

	userRepoMock.On("GetUserByLogin", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	loginForTest := "admin"
	actualResult := ag.IsLoginAvailable(loginForTest)

	expectedResult := true

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsLoginAvailable_ShouldBeNotAvailable(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	//	Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByLogin", mock.Anything).Return(user, nil)

	ag := NewAuthGuard(userRepoMock)

	loginForTest := "admin"
	actualResult := ag.IsLoginAvailable(loginForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsLoginAvailable_ShouldBeNotAvailableEmpty(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	ag := NewAuthGuard(userRepoMock)

	loginForTest := ""
	actualResult := ag.IsLoginAvailable(loginForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsLoginAvailable_Error(t *testing.T) {
	userRepoMock := &mocks.UserRepositoryMock{}

	//	Mocked user to get
	user := domain.User{
		ID:           "5d10c26392ab2a29a9df1ab8",
		Name:         "adminName",
		Surname:      "adminSurname",
		Login:        "admin",
		PasswordHash: "adminadmin",
		Role:         domain.Admin,
	}

	userRepoMock.On("GetUserByLogin", mock.Anything).Return(user, errors.New("someError"))

	ag := NewAuthGuard(userRepoMock)

	loginForTest := "admin"
	actualResult := ag.IsLoginAvailable(loginForTest)

	expectedResult := false

	assert.Equal(t, expectedResult, actualResult)
}
