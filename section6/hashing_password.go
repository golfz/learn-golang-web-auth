package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

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
