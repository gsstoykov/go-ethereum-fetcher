package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/model"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handles user-related requests.
type UserHandler struct {
	ur        repository.IUserRepository
	jwtSecret []byte
}

// JWTClaims represents the JWT claims.
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(ur repository.IUserRepository) *UserHandler {
	return &UserHandler{
		ur:        ur,
		jwtSecret: []byte(os.Getenv("JWT_STRING")),
	}
}

// CreateUser handles the creation of a new user.
// It binds the JSON request to the user model, hashes the password, and stores the user in the database.
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var u model.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		return
	}

	cu, err := uh.ur.Create(&model.User{
		Username: u.Username,
		Password: string(hash),
	})
	if err != nil {
		log.Printf("Failed to create user: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user!"})
		return
	}

	// Respond with the created user
	ctx.JSON(http.StatusCreated, gin.H{"user": cu})
}

// Authenticate handles user authentication.
// It verifies the username and password, generates a JWT token, and returns it to the client.
func (uh *UserHandler) Authenticate(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// Lookup the user in the database
	storedUser, err := uh.ur.FindByUsername(user.Username)
	if err != nil {
		log.Printf("Failed to find user by username: %v\n", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password!"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		log.Printf("Password mismatch for user %s: %v\n", user.Username, err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password!"})
		return
	}

	// Generate JWT token
	token, err := uh.generateToken(storedUser.Username, time.Hour*12)
	if err != nil {
		log.Printf("Failed to generate token for user %s: %v\n", user.Username, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token!"})
		return
	}

	// Log successful authentication
	log.Printf("User %s authenticated successfully\n", user.Username)

	// Set token in the response header
	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// generateToken generates a JWT token for a user.
func (uh *UserHandler) generateToken(username string, expiration time.Duration) (string, error) {
	claims := uh.createJWTClaims(username, expiration)
	return uh.signToken(claims)
}

// createJWTClaims creates JWT claims for a user.
func (uh *UserHandler) createJWTClaims(username string, expiration time.Duration) *JWTClaims {
	return &JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

// signToken signs the JWT claims with the secret.
func (uh *UserHandler) signToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(uh.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// FetchUsers handles fetching all users.
// It retrieves all users from the database and returns them in the response.
func (uh *UserHandler) FetchUsers(ctx *gin.Context) {
	users, err := uh.ur.FindAll()
	if err != nil {
		log.Printf("Failed to fetch users: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// FetchUserTransactions handles fetching transactions for a specific user.
// It retrieves the transactions associated with the authenticated user and returns them in the response.
func (uh *UserHandler) FetchUserTransactions(ctx *gin.Context) {
	username, exists := ctx.Get("username")
	if !exists {
		log.Println("Unauthorized request")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := uh.ur.FindByUsername(fmt.Sprint(username))
	if err != nil {
		log.Printf("Failed to find user by username: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user!"})
		return
	}

	txs, err := uh.ur.FindUserTransactions(user.ID)
	if err != nil {
		log.Printf("Failed to fetch transactions for user %s: %v\n", user.Username, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"usertransactions": txs})
}
