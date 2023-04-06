package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
)

func main() {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// fmt.Println("PublicKey", privKey.PublicKey)
	// fmt.Println(fmt.Printf("priv %+v", privKey))
	fmt.Println("priv", x509.MarshalPKCS1PrivateKey(privKey))
	fmt.Println("pub", len(x509.MarshalPKCS1PublicKey(&privKey.PublicKey)))
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
