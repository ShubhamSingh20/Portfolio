package auth


import (
	//"bytes"
	"fmt"
	"io"
	"os"
	"crypto/rand"
	"crypto/sha1"
)

const saltSize = 16

func generateSalt(secret string) string {
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
	
	salt := fmt.Sprintf("%x", hash.Sum(buf))
	
	return salt
}


func gethashPassword(plainPassword, salt string) string {
	passwordHash := sha1.New() 
	combination := salt + plainPassword
	
	io.WriteString(passwordHash, combination)
	hashedPassword := fmt.Sprintf("%x", passwordHash.Sum(nil))

	return hashedPassword
}

func getHashedPasswordAndSalt(plainPassword string) (string, string) {
	salt := generateSalt(plainPassword)
	passwordHashString := gethashPassword(plainPassword, salt)

	return passwordHashString, salt
}