package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mundaelol/ShareX-Screenshot-Uploader/utils"
)

func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/images/" {
		RootHandler(w, r)
		return
	}

	imageName := strings.TrimPrefix(r.URL.Path, "/images/")
	imagePath := filepath.Join(utils.Directory, imageName)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	file, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Error opening image file: %v", err)
		http.Error(w, "Failed to open image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/"+filepath.Ext(imagePath)[1:])
	http.ServeContent(w, r, imageName, time.Now(), file)
}
