package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mundaelol/ShareX-Screenshot-Uploader/utils"
)

func UploadHandler(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !utils.CheckRateLimit() {
			http.Error(w, "Rate limit exceeded, try again later", http.StatusTooManyRequests)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseMultipartForm(utils.UploadFileSize); err != nil {
			http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}

		if r.FormValue("apiKey") != apiKey {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		file, fileHeader, err := r.FormFile("Image")
		if err != nil {
			http.Error(w, "Failed to retrieve file: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		fileName := utils.SanitizeFileName(fileHeader.Filename)

		targetFile := filepath.Join(utils.Directory, fileName)
		outFile, err := os.Create(targetFile)
		if err != nil {
			log.Printf("Error creating file %s: %v", targetFile, err)
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, file); err != nil {
			log.Printf("Error saving file %s: %v", targetFile, err)
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		uploadedFileURL := utils.DomainURL + fileName
		log.Printf("File uploaded successfully: %s", uploadedFileURL)
		fmt.Fprint(w, uploadedFileURL)
	}
}
