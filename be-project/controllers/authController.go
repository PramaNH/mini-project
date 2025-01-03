package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"be-project/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("A7d!zPp*&$aW9TgMx@1#qLkN")
var queries *db.Queries

func InitDB(q *db.Queries) {
	queries = q
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Debug: Log request body
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("Received login request: %+v\n", creds)

	// Query user dari database
	user, err := queries.GetUserByUsername(context.Background(), creds.Username)
	if err != nil {
		log.Println("Error fetching user from database:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Debug: Log ambil user
	log.Printf("Fetched user from database: %+v\n", user)

	// Validasi password
	if user.Password != creds.Password {
		log.Println("Password mismatch")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error generating token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	log.Println("Generated token:", tokenString)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
