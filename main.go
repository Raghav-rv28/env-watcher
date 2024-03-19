package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

// Directory to watch
var DIRECTORY_TO_WATCH = "/home/raghav/code"

// init function which adds environment variables to runtime.
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Function to upload file to S3
func uploadToS3(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	AWS_REGION, exists := os.LookupEnv("AWS_REGION")
	if !exists {
		fmt.Println("env variable not found")
	}
	AWS_SECRET_KEY, exists := os.LookupEnv("AWS_SECRET_KEY")
	if !exists {
		fmt.Println("env variable not found")
	}
	AWS_ACCESS_KEY_ID, exists := os.LookupEnv("AWS_ACCESS_KEY_ID")
	if !exists {
		fmt.Println("env variable not found")
	}
	AWS_BUCKET_NAME, exists := os.LookupEnv("AWS_BUCKET_NAME")
	if !exists {
		fmt.Println("env variable not found")
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_KEY, ""),
	}))

	svc := s3.New(sess)
	dir := filepath.Base(filepath.Dir(filePath))
	fmt.Println(dir)

	filename := filepath.Base(filePath)
	key := filepath.Join(dir, filename)

	fmt.Println(key)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(AWS_BUCKET_NAME),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
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
			if event.Op&fsnotify.Create == fsnotify.Create {
				if strings.Contains(event.Name, ".env") || strings.Contains(event.Name, ".env.local") {
					fmt.Println("New file created:", event.Name)
					// dir := filepath.Dir(event.Name)
					// // Create a directory with the same name as the file being uploaded
					// errDir := os.MkdirAll(dir, os.ModePerm)
					// if errDir != nil {
					// 	fmt.Println("Error creating directory:", err)
					// }
					// uploadToS3
					err := uploadToS3(event.Name)
					if err != nil {
						fmt.Println("Error uploading to S3:", err)
					} else {
						fmt.Println("File uploaded to S3 successfully.")
					}
				}
			}
		case err := <-watcher.Errors:
			fmt.Println("Error:", err)
		}
	}
}
