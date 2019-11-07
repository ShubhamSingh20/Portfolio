package auth


import (
	"bytes"
	"fmt"
	"io"
	"os"
	"crypto/rand"
	"crypto/sha1"
)

const saltSize = 16

func generateSalt(secret string) []byte {
	secretByte := []byte(secret)

	buf := make([]byte, saltSize, saltSize + sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secretByte)
	
	
	return hash.Sum(buf)
}


func gethashPassword(plainPassword, salt string) []byte {
	passwordHash := sha1.New() 
	combination := salt + plainPassword
	
	io.WriteString(passwordHash, combination)
	return passwordHash.Sum(nil)
}

func getHashedPasswordAndSalt(plainPassword string) (string, string) {
	salt := fmt.Sprintf("%x", generateSalt(plainPassword))
	passwordHashString := fmt.Sprintf("%x", gethashPassword(plainPassword, salt)) 

	return passwordHashString, salt
	
}


func authenticate(inputPassword, correctHashPassword, salt string) bool {
	inputPasswordHash := gethashPassword(inputPassword, salt)
	
	return bytes.Equal(
		[]byte(correctHashPassword), []byte(fmt.Sprintf("%x",inputPasswordHash)),
	)
}