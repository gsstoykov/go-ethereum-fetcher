package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	mocks "github.com/gsstoykov/go-ethereum-fetcher/fetcher/mock"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
)

func TestCreateUser(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	handler := NewUserHandler(mockUserRepo)

	// Sample user data
	user := model.User{
		Username: "testuser",
		Password: "password", // Raw password
	}

	// Generate a consistent hash for the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	expectedUser := model.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

	// Setup mock to expect any user with a matching username
	mockUserRepo.On("Create", mock.AnythingOfType("*model.User")).Return(&expectedUser, nil).Run(func(args mock.Arguments) {
		// Assert on the argument here if necessary
		createdUser := args.Get(0).(*model.User)
		assert.Equal(t, user.Username, createdUser.Username)
		// Optionally, check for password if required
		assert.NotEqual(t, user.Password, createdUser.Password) // Ensures the password is hashed
	})

	// Setup Gin router and request
	router := gin.Default()
	router.POST("/users", handler.CreateUser)

	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), user.Username)

	// Assert mock expectations
	mockUserRepo.AssertExpectations(t)
}

func TestFetchUsers(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	handler := NewUserHandler(mockUserRepo)

	// Sample users
	users := []model.User{
		{Username: "testuser1"},
		{Username: "testuser2"},
	}

	// Setup mock
	mockUserRepo.On("FindAll").Return(users, nil)

	// Setup Gin router and request
	router := gin.Default()
	router.GET("/users", handler.FetchUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "testuser1")
	assert.Contains(t, w.Body.String(), "testuser2")

	mockUserRepo.AssertExpectations(t)
}
