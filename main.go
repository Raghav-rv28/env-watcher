package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fsnotify/fsnotify"
)

// vars to control s3 login creds and `.env` file location
var (
	DIRECTORY_TO_WATCH = os.Getenv("DIRECTORY_TO_WATCH")
	AWS_ACCESS_KEY_ID  = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_KEY     = os.Getenv("AWS_SECRET_KEY")
	AWS_REGION         = os.Getenv("AWS_REGION")
	AWS_BUCKET_NAME    = os.Getenv("AWS_BUCKET_NAME")
)

// Function to upload file to S3
func uploadToS3(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

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
	fileError := filepath.Walk(DIRECTORY_TO_WATCH,
		func(path string, info os.FileInfo, err error) error {
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

	if fileError != nil {
		fmt.Fscanf(os.Stderr, "%v", fileError)
		return
	}

	// Process events
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
				if (strings.Contains(event.Name, ".env") || strings.Contains(event.Name, ".env.local")) && !strings.Contains(event.Name, "~") {
					fmt.Println("New file event:", event.Name)
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
