package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {
	privKey, err := rsa.GenerateKey(rand.Reader, 1372)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("priv", privKey.PublicKey.N)

	message := "secret message"

	encryptedMSG, err := rsa.EncryptPKCS1v15(rand.Reader, &privKey.PublicKey, []byte(message))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("encryptedMSG", string(encryptedMSG))
	decryptedMSG, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, encryptedMSG)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("encryptedMSG", string(decryptedMSG))
}
