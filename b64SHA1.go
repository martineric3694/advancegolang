package main

import (
	"crypto/des"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

func encodeDecodeBase64(data string) {

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)
}

func encodeSHA1(data string) {
	var text = data
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
}

func doHashUsingSalt(data string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", data, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func Hashing(message, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func tripleDES(message string) []byte {
	ede2Key := []byte(message)

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	fmt.Println(tripleDESKey)
	_, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}
	return tripleDESKey
}
