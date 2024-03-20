package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// FIXME: Enter your directory to watch
// Key used for AES encryption (must be 16, 24, or 32 bytes long)
var (
	DIRECTORY_TO_WATCH = "/home/raghav/code"
	encryptionKey      = []byte("")
)

// Function to encrypt a file
func encryptFile(filePath string) error {
	// Open the input file
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFilePath := filePath + ".enc"
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

func main() {
	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	// Function to walk through directory tree and watch for changes
	filepath.Walk(DIRECTORY_TO_WATCH, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
		}
		return nil
	})

	// Process events
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
				if (strings.Contains(event.Name, ".env") || strings.Contains(event.Name, ".env.local")) &&
					!strings.Contains(event.Name, "~") && !strings.Contains(event.Name, ".enc") {
					fmt.Println("New file event:", event.Name, event.Op)
					// uploadToS3
					err := encryptFile(event.Name)
					if err != nil {
						fmt.Println("Error Encrypting the file:", err)
					} else {
						fmt.Println("File successfully Encrypted and stored.")
					}
				}
			}
		case err := <-watcher.Errors:
			fmt.Println("Error:", err)
		}
	}
}
