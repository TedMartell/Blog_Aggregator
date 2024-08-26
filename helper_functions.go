package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func generateAPI() (string, error) {

	buf := make([]byte, 256)

	_, err := rand.Read(buf)
	if err != nil {
		log.Printf("error generating api_key: %v", err)
		return "", err
	}
	foo := sha256.New()
	foo.Write(buf)
	hash := foo.Sum(nil)
	encoded := hex.EncodeToString(hash)

	return encoded, nil
}
