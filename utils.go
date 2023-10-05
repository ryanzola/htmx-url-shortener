package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Env{
		Db:       os.Getenv("DB"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
	}
}

func RanHash() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var slug string

	for i := 0; i < 8; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			fmt.Println("Error generating random index:", err)
			return ""
		}
		slug += string(characters[randomIndex.Int64()])
	}

	return slug
}
