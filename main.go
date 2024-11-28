package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mundae-uploader/handlers"
	"mundae-uploader/utils"

	"github.com/joho/godotenv"
)

var apiKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	apiKey = os.Getenv("UPLOAD_API_KEY")
	if apiKey == "" {
		log.Fatal("UPLOAD_API_KEY environment variable is not set")
	}

	utils.EnsureDir()
}

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/images/", handlers.ImagesHandler)
	http.HandleFunc("/upload/screenshot", handlers.UploadHandler(apiKey))

	fmt.Print("\033[H\033[2J")
	log.Printf("Server started at %s%s", utils.Host, utils.Port)
	log.Fatal(http.ListenAndServe(utils.Port, nil))
}
