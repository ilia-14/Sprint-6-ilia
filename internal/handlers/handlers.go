package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error getting the file:", err)
		http.Error(w, "Error getting file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	result, err := service.AutoConvert(string(content))
	if err != nil {
		log.Println("Error during conversion:", err)
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		log.Println("Error creating uploads directory:", err)
		http.Error(w, "Error creating uploads directory", http.StatusInternalServerError)
		return
	}

	fileName := filepath.Join("uploads", time.Now().UTC().Format("2006-01-02_15-04-05")+".txt")
	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		log.Println("Error writing to file:", err)
		http.Error(w, "Error saving the conversion result", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	res, err := w.Write([]byte(result))
	if err != nil {
		log.Println("Error writing response:", err)
		http.Error(w, "Error sending response", http.StatusInternalServerError)
		return
	}

	log.Printf("Sent %d bytes of data to client", res)
}
