package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"
)

func encryptFile(filePath string, encryptionKey []byte) error {
	// Open the input file
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFilePath := filePath + ".share.enc"
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Create AES cipher block using the encryption key
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return err
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Write the IV to the output file
	if _, err := outputFile.Write(iv); err != nil {
		return err
	}

	// Create the AES cipher in Galois Counter Mode (GCM) with the given IV
	stream := cipher.NewCFBEncrypter(block, iv)

	// Create a writer that encrypts data as it is written to the output file
	writer := &cipher.StreamWriter{S: stream, W: outputFile}

	// Copy data from the input file to the encrypted output file
	if _, err := io.Copy(writer, inputFile); err != nil {
		return err
	}

	fmt.Printf("File encrypted successfully: %s\n", outputFilePath)

	return nil
}

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
	// arguments
	if len(os.Args) < 2 {
		fmt.Println("Enter the arguments !Usage: cryptor <'encrypt' | 'decrypt'> <file> <secret-key>")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Error: file Path not found, Usage: cryptor <'encrypt' | 'decrypt'> <file> <secret-key>")
		return
	}
	if len(os.Args) < 4 {
		fmt.Println("ENCRYPTION_KEY not found, Usage: cryptor <'encrypt' | 'decrypt'> <file> <secret-key>")
		return
	}
	process := os.Args[1]
	filePath := os.Args[2]
	encryptionKey := os.Args[3]
	// Decrypt the file
	if strings.Contains(process, "decrypt") && strings.Contains(filePath, ".enc") {
		err := decryptFile(filePath, []byte(encryptionKey))
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		// Encrypt the file
	} else if strings.Contains(process, "encrypt") && !strings.HasSuffix(filePath, ".enc") {
		fmt.Println("Encrypting the file", filePath, "using the following key:", encryptionKey)
		err := encryptFile(filePath, []byte(encryptionKey))
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}

	}
}
