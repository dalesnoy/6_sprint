package handlers

import (
	"bytes"
	//"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("./index.html"); os.IsNotExist(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./index.html")

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не распознан", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")

	if err != nil {
		http.Error(w, "Ошибка доступа к файлу: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	result, err := service.Converter(bytes.NewReader(data))

	if err != nil {
		http.Error(w, "Преобразование не выполнено", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)

	if ext == "" {
		ext = ".txt"
	}

	name := time.Now().UTC().Format("20060102_150405") + ext
	path := filepath.Join(".", name)

	if err := os.WriteFile(path, []byte(result), 0644); err != nil {
		http.Error(w, "Ошика записи файла:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(result))

}
