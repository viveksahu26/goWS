package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

var (
	baseURL = os.Getenv("BASE_URL")
	port    = os.Getenv("PORT")

	showTmpl  *template.Template
	indexTmpl *template.Template
	errorTmpl *template.Template
)

func init() {
	if baseURL == "" {
		baseURL = "http://localhost:8888"
	}
	if port == "" {
		port = "8888"
	}

	// Load templates
	var err error
	indexTmpl, err = template.ParseFiles("tmpl/index.html")
	if err != nil {
		log.Fatal("Failed to parse index.html:", err)
	}
	showTmpl, err = template.ParseFiles("tmpl/show.html")
	if err != nil {
		log.Fatal("Failed to parse show.html:", err)
	}
	errorTmpl, err = template.ParseFiles("tmpl/error.html")
	if err != nil {
		log.Fatal("Failed to parse error.html:", err)
	}
}

func main() {
	http.HandleFunc("/qr/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		embedURL := r.PostFormValue("url")
		sizeStr := r.PostFormValue("size")
		size, err := strconv.Atoi(sizeStr)
		if err != nil || (size != 256 && size != 512) {
			size = 256 // Default to small
		}

		// Validate URL
		if _, err := url.ParseRequestURI(embedURL); err != nil {
			log.Printf("Invalid URL submitted: %s", embedURL)
			renderError(w, "Invalid URL provided. Please enter a valid URL (e.g., https://example.com).")
			return
		}

		embedURLB64 := base64.URLEncoding.EncodeToString([]byte(embedURL))
		http.Redirect(w, r, fmt.Sprintf("/qr/show/%s?size=%d", embedURLB64, size), http.StatusFound)
	})

	http.HandleFunc("/qr/view/", func(w http.ResponseWriter, r *http.Request) {
		embedURLB64 := r.URL.Path[len("/qr/view/"):]
		embedURL, err := base64.URLEncoding.DecodeString(embedURLB64)
		if err != nil {
			http.Error(w, "Invalid QR code URL", http.StatusBadRequest)
			return
		}

		sizeStr := r.URL.Query().Get("size")
		size, err := strconv.Atoi(sizeStr)
		if err != nil || (size != 256 && size != 512) {
			size = 256
		}

		qr, err := qrcode.Encode(string(embedURL), qrcode.Low, size)
		if err != nil {
			log.Printf("Failed to generate QR: %v", err)
			http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		if _, err := w.Write(qr); err != nil {
			log.Printf("Failed to write QR image: %v", err)
		}
	})

	http.HandleFunc("/qr/show/", func(w http.ResponseWriter, r *http.Request) {
		embedURLB64 := r.URL.Path[len("/qr/show/"):]
		embedURL, err := base64.URLEncoding.DecodeString(embedURLB64)
		if err != nil {
			renderError(w, "Invalid QR code URL")
			return
		}

		sizeStr := r.URL.Query().Get("size")
		size, err := strconv.Atoi(sizeStr)
		if err != nil || (size != 256 && size != 512) {
			size = 256
		}

		w.Header().Set("Content-Type", "text/html")
		qrImageURL := fmt.Sprintf("%s/qr/view/%s?size=%d", baseURL, embedURLB64, size)
		data := struct {
			URL        string
			QRImageURL string
		}{
			URL:        string(embedURL),
			QRImageURL: qrImageURL,
		}

		if err := showTmpl.Execute(w, data); err != nil {
			log.Printf("Failed to render show template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			renderError(w, "Page not found")
			return
		}
		w.Header().Set("Content-Type", "text/html")
		if err := indexTmpl.Execute(w, nil); err != nil {
			log.Printf("Failed to render index template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func renderError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html")
	data := struct{ Message string }{Message: message}
	if err := errorTmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
