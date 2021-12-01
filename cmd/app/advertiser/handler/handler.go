package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/database"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/database/model"
)

var (
	landingPagePath   = os.Getenv("LANDING_PAGE_DIR_PATH")
	conversionDirPath = os.Getenv("CV_PAGE_DIR_PATH")
	Advertiser        = os.Getenv("ADVERTISER_DOMAIN")
	AdDeliver         = os.Getenv("AD_DELIVER_DOMAIN")
	Publisher         = os.Getenv("PUBLISHER_DOMAIN")
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles(landingPagePath)
	if err != nil {
		log.Printf("Invalid template file path: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, Advertiser); err != nil {
		log.Printf("Failed return template file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Success return template file")
}

func CVPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	file := filepath.Base(r.URL.Path)

	t, err := template.ParseFiles(conversionDirPath + file)
	if err != nil {
		log.Printf("Invalid template file path: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, struct {
		AdDeliver  string
		Advertiser string
		Publisher  string
	}{
		AdDeliver:  Advertiser,
		Advertiser: Advertiser,
		Publisher:  Publisher,
	}); err != nil {
		log.Printf("Failed return template file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Success return template file")
}

func Report(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Invalid request method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var report *model.Report
	if err := json.NewDecoder(r.Body).Decode(report); err != nil {
		log.Printf("Invalid request body format: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := database.InsertReport(report, r.Referer(), r.Host); err != nil {
		log.Printf("Failed insert report to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Success receive report")
	w.WriteHeader(http.StatusOK)
}
