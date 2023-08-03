package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/joho/godotenv"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ConversionResponse struct {
	FileName string `json:"fileName"`
	DataURL  string `json:"dataURL"`
}

func convertToWebP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error reading image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Error decoding image", http.StatusBadRequest)
		return
	}

	// Получаем значение степени сжатия из переменной окружения COMPRESS_RATIO
	compressRatioStr := os.Getenv("COMPRESS_RATIO")
	compressRatio, err := strconv.Atoi(compressRatioStr)
	if err != nil || compressRatio <= 0 {
		compressRatio = 80 // Значение по умолчанию, если переменная окружения не задана или имеет некорректное значение.
	}

	webpBuf := new(bytes.Buffer)
	if err := webp.Encode(webpBuf, img, &webp.Options{Lossless: false, Quality: float32(compressRatio)}); err != nil {
		http.Error(w, "Error converting to webp", http.StatusInternalServerError)
		return
	}

	// Получаем имя исходного файла без расширения
	fileName := strings.TrimSuffix(header.Filename, ".png")
	fileName = strings.TrimSuffix(fileName, ".jpg")
	fileName = strings.TrimSuffix(fileName, ".jpeg")

	// Формируем ответ в формате JSON
	response := ConversionResponse{
		FileName: fileName + ".webp",
		DataURL:  fmt.Sprintf("data:image/webp;base64,%s", base64.StdEncoding.EncodeToString(webpBuf.Bytes())),
	}

	// Отправляем ответ в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Выводим информацию в лог
	fileSize := len(webpBuf.Bytes())
	clientIP := r.RemoteAddr
	duration := time.Since(startTime)
	log.Printf("Time: %s, File: %s, Size: %d bytes, IP: %s, Result: Success, Duration: %s", startTime.Format("2006-01-02 15:04:05"), fileName, fileSize, clientIP, duration)
}

func serveUI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file. Using default values.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию, если переменная окружения не задана.
	}

	compressRatioStr := os.Getenv("COMPRESS_RATIO")
	compressRatio, err := strconv.Atoi(compressRatioStr)
	if err != nil || compressRatio <= 0 {
		compressRatio = 80 // Значение по умолчанию, если переменная окружения не задана или имеет некорректное значение.
	}

	// Выводим значения переменных в консоль
	log.Printf("PORT=%s\n", port)
	log.Printf("COMPRESS_RATIO=%d\n", compressRatio)

	http.HandleFunc("/convert", convertToWebP)
	http.HandleFunc("/ui", serveUI)

	addr := ":" + port
	log.Printf("Server started on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
