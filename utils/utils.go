package utils

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	UploadFileSize = 500 * 1024 * 1024 // 500 MB
	FileNameLength = 12                // Length of randomized file names
	RateLimit      = 10                // Maximum requests per second
)

var Host = "http://localhost"
var Port = ":8080"
var Directory = "./images/"
var DomainURL = Host + Port + "/images/" // Base URL to access the images

var (
	randSrc       = rand.New(rand.NewSource(time.Now().UnixNano()))
	rateLimitChan = make(chan struct{}, RateLimit)
)

func init() {
	for i := 0; i < RateLimit; i++ {
		rateLimitChan <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case rateLimitChan <- struct{}{}:
			default:
			}
		}
	}()
}

func EnsureDir() {
	if _, err := os.Stat(Directory); os.IsNotExist(err) {
		if err := os.MkdirAll(Directory, os.ModePerm); err != nil {
			log.Fatalf("Failed to create sharex directory: %v", err)
		}
	}
}

func SanitizeFileName(fileName string) string {
	ext := filepath.Ext(filepath.Base(fileName))
	return RandomString(FileNameLength) + ext
}

func RandomString(length int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[randSrc.Intn(len(chars))]
	}
	return string(result)
}

func CheckRateLimit() bool {
	select {
	case <-rateLimitChan:
		return true
	default:
		return false
	}
}
