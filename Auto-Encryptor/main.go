package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func containsSubstring(s string, substrings []string) bool {
	for _, sub := range substrings {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

func readEnvFile(fileLocation string) (map[string]string, error) {
	// Open the file
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a map to store environment variables
	envVars := make(map[string]string)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line by '='
		parts := strings.Split(scanner.Text(), "=")
		// Store the environment variable and its value in the map
		envVars[parts[0]] = parts[1]
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envVars, nil
}

// Function to encrypt a file
func encryptFile(filePath string, encryptionKey []byte) error {
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
	var encryptionKey string
	var watchDirectory string
	ignoreFolders := []string{".", "node_modules"}

	// get current user to retrieve the env file.
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	envVars, err := readEnvFile("/home/" + currentUser.Username + "/.file_watcher_env")
	if err != nil {
		panic(err)
	}
	fmt.Print(envVars)
	// Retrieve encryptionKey, directory to watch and ignored directory.
	encryptionKey = envVars["encryption_key"]
	watchDirectory = envVars["watch_directory"]
	ignoreDirManual := strings.Split(envVars["ignore_dir"], ";")
	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	// Function to walk through directory tree and watch for changes
	filepath.Walk(watchDirectory, func(path string, info os.FileInfo, err error) error {
		dirName := filepath.Base(path)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if info.IsDir() && containsSubstring(dirName, append(ignoreDirManual, ignoreFolders...)) {
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
					!strings.Contains(event.Name, "~") && !strings.Contains(event.Name, ".enc") && !strings.Contains(event.Name, ".share") {
					fmt.Println("New file event:", event.Name, event.Op)
					// uploadToS3
					err := encryptFile(event.Name, []byte(encryptionKey))
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
