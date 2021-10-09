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
	fmt.Println(nonceSize)
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}

func encryptFile(sourceFile, passphrase string) {

	data, _ := ioutil.ReadFile(sourceFile)

	file, _ := os.Create("e" + sourceFile)
	defer file.Close()

	file.Write(encrypt(data, passphrase))
}

func decryptFile(fileName, passphrase string) {
	data, _ := ioutil.ReadFile(fileName)
	fmt.Println("len data: ", len(data))

	decryptData := decrypt(data, passphrase)
	ioutil.WriteFile("d"+fileName, decryptData, 0)
}

func main() {
	sourceFile := "test.txt"

	/*
		ciphertext := encrypt([]byte("hello worlds"), mypass)
		fmt.Println(string(ciphertext))
		fmt.Println()

		plaintext := decrypt(ciphertext, mypass)
		fmt.Println(string(plaintext))
	*/
	encryptFile(sourceFile, mypass)
	decryptFile("e"+sourceFile, mypass)
	fmt.Println("Done")

}
