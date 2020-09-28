package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionId int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionId == 0 {
		return fmt.Errorf("Invalid session ID")
	}

	return nil
}

var Key = []byte{}

func main() {
	pass := "12345678"

	hashedPass := hashPassword(pass)

	fmt.Println("this is a pass:", pass)
	fmt.Println("this is a hashed:", string(hashedPass))

	err := comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatal("You not logged in")
	}

	fmt.Println("You logged in")

	for i := 1; i <= 64; i++ {
		Key = append(Key, byte(i))
	}

	message := []byte("Hi, I am GolFz")

	signature, err := signMessage(message)
	if err != nil {
		log.Println("sign message error:", err)
	}

	log.Println("signature:", string(signature))

	isValid, err := checkSignature(message, signature)
	if err != nil {
		log.Fatal("checkSignature error", err)
	}

	if isValid {
		log.Println("signature is valid")
	} else {
		log.Println("signature is invalid")
	}
}

func hashPassword(password string) []byte {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return bs
}

func comparePassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("Error when compare password: %w", err)
	}
	return nil
}

func signMessage(message []byte) ([]byte, error) {
	h := hmac.New(sha512.New, Key)
	_, err := h.Write(message)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while write hmac: %w", err)
	}
	hashed := h.Sum(nil)
	return hashed, nil
}

func checkSignature(message, signature []byte) (bool, error) {
	expectedSignature, err := signMessage(message)
	if err != nil {
		return false, fmt.Errorf("Error in checkSignature while sign message")
	}

	isValid := hmac.Equal(expectedSignature, signature)
	return isValid, nil
}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(Key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		return Key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}
