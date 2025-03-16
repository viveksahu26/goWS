package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh/terminal"
)

func decryptECDSAKey(filePath string, password []byte) ([]byte, error) {
	// Read the encrypted key from the file
	encryptedKey, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read encrypted key file: %v", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(encryptedKey)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Decrypt the PKCS#8 encrypted private key
	decryptedKey, err := x509.DecryptPKCS8PrivateKey(block.Bytes, password)
	x509.DecryptPEMBlock()
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt key: %v", err)
	}

	// Convert the decrypted key into a PEM block
	// Note: This step depends on the type of key (RSA, ECDSA, etc.) and may vary
	// Here, we're assuming an ECDSA key for demonstration
	ecdsaPrivateKey, ok := decryptedKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("decrypted key is not of type ECDSA")
	}

	// Marshal the ECDSA private key into DER format
	derKey, err := x509.MarshalECPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ECDSA private key: %v", err)
	}

	// Encode the DER key into a PEM block
	pemBlock := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derKey,
	}

	var pemData []byte
	pemData = append(pemData, pem.EncodeToMemory(&pemBlock)...)

	return pemData, nil
}

func main() {
	fmt.Println("Enter the decryption password:")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("Failed to read password:", err)
		return
	}
	fmt.Println()

	// Specify the path to your encrypted key file
	encryptedKeyPath := "path/to/your/encrypted/key.pem"
	decryptedKey, err := decryptECDSAKey(encryptedKeyPath, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Decrypted key:")
	fmt.Println(string(decryptedKey))
}
