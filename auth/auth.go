package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

// Secret key used for signing the JWT token (keep it safe and secure)
var jwtKey = []byte("secret_key_12345")

// Claims struct to be encoded in JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to generate a new JWT token
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Create the JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Middleware to authenticate JWT token
func AuthenticateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		// Extract the JWT token from the Authorization header
		tokenStr := strings.Split(authHeader, " ")[1]
		claims := &Claims{}

		// Verify and parse the JWT token
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		// Check if the token is valid
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// Check if the token is expired
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// If the token is valid, call the next handler
		next.ServeHTTP(w, r)
	}
}

// Function to handle login and generate JWT token
func Login(w http.ResponseWriter, r *http.Request) {

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// Decode the request payload
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Check if the credentials are valid
	if creds.Username != "admin" || creds.Password != "password" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// Generate a new JWT token
	token, err := GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
