package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var mypass = "mypass"

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	//ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return gcm.Seal(nonce, nonce, data, nil)

}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}

func encryptFile(data []byte, fileName, passphrase string) {
	file, _ := os.Create(fileName)
	defer file.Close()
	file.Write(encrypt(data, passphrase))
}

func decryptFile(fileName, passphrase string) []byte {
	data, _ := ioutil.ReadFile(fileName)
	return decrypt(data, passphrase)
}

func main() {
	ciphertext := encrypt([]byte("hello worlds"), mypass)
	fmt.Println(string(ciphertext))
	fmt.Println()

	plaintext := decrypt(ciphertext, mypass)
	fmt.Println(string(plaintext))

	fmt.Println()
	encryptFile([]byte("hello world"), "example.txt", mypass)
	plaintext = decryptFile("example.txt", mypass)
	fmt.Println(string(plaintext))

}
