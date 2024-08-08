package infrastructure

import (
	"os"
	domain "task-management-api-mongodb/Domain"

	jwt "github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateToken(u domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		UserID:  u.ID,
		Email:   u.Email,
		IsAdmin: u.IsAdmin,
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func ValidateToken(jwtToken string) (*customClaims, bool) {
	claims := &customClaims{}
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Vaild {
		return claims, false
	}
	return claims, true
}
