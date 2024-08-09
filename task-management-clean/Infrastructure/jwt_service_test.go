package infrastructure

import (
	domain "clean-architecture/Domain"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwtSecret := "testsecret"
	t.Setenv("JWT_SECRET", jwtSecret)
	tests := []domain.User{
		{
			ID:      "sdjfksdj",
			Email:   "email@email.com",
			IsAdmin: true,
		},
		{
			ID:      "sdjsadfsdfds",
			Email:   "email@sdds.com",
			IsAdmin: false,
		},
		{
			ID:      "ssdfds",
			Email:   "email@sdsd.com",
			IsAdmin: false,
		},
	}
	for _, tt := range tests {
		token, err := GenerateToken(tt)
		if err != nil {
			t.Fatal(err)
		}
		claims, valid := ValidateToken(token)
		if !valid {
			t.Fatal("expected valid token but found invalid")
		}
		if claims.UserID != tt.ID {
			t.Fatalf("expected id: %v found id: %v ", tt.ID, claims.UserID)
		}
		if claims.Email != tt.Email {
			t.Fatalf("expected email: %v found email: %v ", tt.Email, claims.Email)
		}
		if claims.IsAdmin != tt.IsAdmin {
			t.Fatalf("expected admin role to be %v, but found %v", tt.IsAdmin, claims.IsAdmin)
		}
	}
}

func TestValidateToken(t *testing.T) {
	jwtSecret := "validSecret"
	t.Setenv("JWT_SECRET", jwtSecret)
	users := []domain.User{
		{
			ID:      "sdjfksdj",
			Email:   "email@email.com",
			IsAdmin: true,
		},
		{
			ID:      "sdjsadfsdfds",
			Email:   "email@sdds.com",
			IsAdmin: false,
		},
		{
			ID:      "ssdfds",
			Email:   "email@sdsd.com",
			IsAdmin: false,
		},
	}
	tests := []string{}
	for _, u := range users {
		token, err := GenerateToken(u)
		if err != nil {
			t.Fatal(err)
		}
		tests = append(tests, token)
	}
	for _, tt := range tests {
		_, valid := ValidateToken(tt)
		if !valid {
			t.Fatal("expected the token to be valid but found invalid")
		}
	}
	jwtSecret = "invalidSecret"
	t.Setenv("JWT_SECRET", jwtSecret)
	for _, tt := range tests {
		_, valid := ValidateToken(tt)
		if valid {
			t.Fatal("expected the token to be invalid but found valid")
		}
	}
}
