package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/model"
	"github.com/gsstoykov/go-ethereum-fetcher/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	ur        repository.IUserRepository
	jwtSecret []byte
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewUserHandler(ur repository.IUserRepository) *UserHandler {
	return &UserHandler{
		ur:        ur,
		jwtSecret: []byte(os.Getenv("JWT_STRING")),
	}
}

func (uh UserHandler) CreateUser(ctx *gin.Context) {
	var u model.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cu, err := uh.ur.Create(&model.User{
		Username: u.Username,
		Password: string(hash),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": cu})
}

func (u *UserHandler) Authenticate(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// Lookup the user in the database
	storedUser, err := u.ur.FindByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := u.generateToken(storedUser.Username, time.Hour*12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set token in the response header
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserHandler) generateToken(username string, expiration time.Duration) (string, error) {
	// Create JWT claims
	claims := u.createJWTClaims(username, expiration)

	// Generate token
	tokenString, err := u.signToken(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *UserHandler) createJWTClaims(username string, expiration time.Duration) *JWTClaims {
	// Create standard claims
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiration).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	// Create custom claims
	claims := &JWTClaims{
		Username:       username,
		StandardClaims: standardClaims,
	}

	return claims
}

func (u *UserHandler) signToken(claims *JWTClaims) (string, error) {
	// Generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the JWT secret
	tokenString, err := token.SignedString(u.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uh UserHandler) FetchUsers(ctx *gin.Context) {
	var us []model.User
	us, err := uh.ur.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": us})
}
