package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

const SALT = 13

func GenerateSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	return salt
}

func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

// VALIDAR COINCIDENCIA
func DoPasswordsMatch(hashedPassword, password string, salt []byte) bool {
	var currPasswordHash = HashPassword(password, salt)
	return hashedPassword == currPasswordHash
}
