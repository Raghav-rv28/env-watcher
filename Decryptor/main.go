package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"strings"
)

// Function to decrypt a file
func decryptFile(filePath string, encryptionKey []byte) error {
	// Open the encrypted file
	encryptedFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	// Create the output file
	outputFilePath := strings.TrimSuffix(filePath, ".enc")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Read the IV from the beginning of the encrypted file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(encryptedFile, iv); err != nil {
		return err
	}

	// Create AES cipher block using the encryption key
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return err
	}

	// Create the AES cipher in Galois Counter Mode (GCM) with the given IV
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create a reader that decrypts data as it is read from the encrypted file
	reader := &cipher.StreamReader{S: stream, R: encryptedFile}

	// Copy data from the encrypted file to the output file
	if _, err := io.Copy(outputFile, reader); err != nil {
		return err
	}

	fmt.Printf("File decrypted successfully: %s\n", outputFilePath)

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("ENCRYPTION_KEY not found, Usage: decrypt <file> <secret-key>")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Enter the arguments !Usage: decrypt <file> <secret-key>")
		return
	}

	filePath := os.Args[1]
	encryptionKey := os.Args[2]

	// Decrypt the file if its name contains a certain string
	if strings.Contains(filePath, "env") && strings.HasSuffix(filePath, ".enc") {
		err := decryptFile(filePath, []byte(encryptionKey))
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
	} else {
		fmt.Println("File does not match decryption criteria.")
	}
}
