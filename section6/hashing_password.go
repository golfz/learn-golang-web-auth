package main

import (
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"crypto/hmac"
)

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

	for i := 1; i<=64; i++ {
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
