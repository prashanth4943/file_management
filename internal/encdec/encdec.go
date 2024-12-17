package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"net/url"
)

func EncDecValues(secureType, value string) string {
	// Key for AES (it must be 32 bytes for AES-256)
	key := "d41d8cd98f00b204e9810998ecf85373" // Default key, change as needed

	// Convert key to 32 bytes (AES-256 requires a 32-byte key)
	// This assumes your key is exactly 16 bytes. If it is longer or shorter, you might need to adjust it.
	keyBytes := []byte(key)

	// Use the key to generate an IV (Initialization Vector) (use a constant or generate dynamically)
	iv := keyBytes // This is just for simplicity, in real cases, use a random IV

	// Create AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %v", err)
	}

	if secureType == "enc" {
		// For encryption, ensure the value is padded to match block size
		plaintext := []byte(value)
		blockSize := block.BlockSize()
		padding := blockSize - len(plaintext)%blockSize
		padText := append(plaintext, make([]byte, padding)...)

		// CBC mode
		mode := cipher.NewCBCEncrypter(block, iv)
		ciphertext := make([]byte, len(padText))
		mode.CryptBlocks(ciphertext, padText)

		// Base64 encode the encrypted bytes
		encValue := base64.StdEncoding.EncodeToString(ciphertext)
		return url.QueryEscape(encValue) // URL encode the base64 string
	}

	if secureType == "dec" {
		// For decryption, decode the base64 string
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			log.Fatalf("Error decoding base64: %v", err)
		}

		// CBC mode
		mode := cipher.NewCBCDecrypter(block, iv)
		plaintext := make([]byte, len(decoded))
		mode.CryptBlocks(plaintext, decoded)

		// Unpad the plaintext
		plaintext = plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]

		return string(plaintext)
	}

	return value
}

// // func main() {
// 	// Test encryption and decryption
// 	encValue := EncryptDecryptValues("enc", "Hello World!")
// 	fmt.Printf("Encrypted: %s\n", encValue)

// 	// decValue := EncryptDecryptValues("dec", encValue)
// 	// fmt.Printf("Decrypted: %s\n", decValue)
// // }
