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
	filePath, err := filepath.Abs("index.html")
	if err != nil {
		log.Println("Error locating file:", err)
		http.Error(w, "Error locating file", http.StatusInternalServerError)
		return
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error loading page:", err)
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
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
	fileName := filepath.Join("uploads", time.Now().UTC().Format("2006-01-02_15-04-05")+".txt")
	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		log.Println("Error writing to file:", err)
		http.Error(w, "Error saving the conversion result", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result + "\nConversion result also saved in: " + fileName))
}
