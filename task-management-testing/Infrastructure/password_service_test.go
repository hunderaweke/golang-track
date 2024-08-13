package infrastructure

import (
	"crypto/rand"
	"math/big"
	"testing"
)

const (
	passwordLength = 8
	charSet        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=<>?"
)

func generateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	charSetLength := big.NewInt(int64(len(charSet)))
	for i := range length {
		randomInt, err := rand.Int(rand.Reader, charSetLength)
		if err != nil {
			return "", err
		}
		password[i] = charSet[randomInt.Int64()]
	}
	return string(password), nil
}

func TestPasswordService(t *testing.T) {
	passwords := []string{}
	for range 10 {
		randPass, err := generateRandomPassword(passwordLength)
		if err != nil {
			t.Fatal(err)
		}
		passwords = append(passwords, randPass)
	}
	tests := []string{}
	for _, p := range passwords {
		hashPassword, err := HashPassword(p)
		if err != nil {
			t.Fatal(err)
		}
		tests = append(tests, hashPassword)
	}
	for i, tt := range tests {
		if !ComparePassword(passwords[i], tt) {
			t.Fatal("expected the hash and password to match but found differing")
		}
	}
	for i, tt := range tests {
		if ComparePassword(passwords[i]+"jkj", tt) {
			t.Fatal("expected the hash and password to differ but found matching")
		}
	}
}
