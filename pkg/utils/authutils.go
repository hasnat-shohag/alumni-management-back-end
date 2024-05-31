package utils

import (
	"alumni-management-server/pkg/config"
	"alumni-management-server/pkg/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// GetPasswordHash returns the hashed password.
func GetPasswordHash(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

// CheckPassword checks if the password is correct.
func CheckPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

type CustomClaims struct {
	StudentId string
	Email     string
	Role      string `json:"role"`
	jwt.StandardClaims
}

// GetJwtForUser returns the JWT for the user using studentId.
func GetJwtForUser(user *models.UserDetail) (string, error) {
	now := time.Now().UTC()
	ttl := time.Minute * time.Duration(config.LocalConfig.JwtExpireMinutes)
	claims := CustomClaims{
		StudentId: user.StudentId,
		Email:     user.Email,
		Role:      user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			Subject:   user.Email,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.LocalConfig.JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
