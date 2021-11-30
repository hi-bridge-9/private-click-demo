package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kyu-takahahsi/private-click-demo/cmd/app/publisher/blind_sign"
	"github.com/kyu-takahahsi/private-click-demo/cmd/app/publisher/public_token"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/database"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/database/model"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/validation"
)

var (
	topPagePath = os.Getenv("TOP_PAGE_FILE_PATH")
	beaconPath  = os.Getenv("BEACON_FILE_PATH")
	AdDeliver   = os.Getenv("AD_DELIVER_DOMAIN")
	wellKnown   = "/.well-known/private-click-measurement"
)

func TopPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles(topPagePath)
	if err != nil {
		log.Printf("Ivalid template file path: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, AdDeliver); err != nil {
		log.Printf("Failed return template file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Success return html template")
}

func TriggerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	param := strings.TrimPrefix(r.URL.Path, "/cv-trigger/")
	li := strings.Split(param, "/")

	if len(li) == 0 || len(li) > 2 {
		log.Printf("Ivalid parameter length: %d", len(li))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var location string
	if len(li) == 2 {
		if validation.IsValidTriggerDataAndPriority(li[0], li[1]) {
			location = fmt.Sprintf("http://%s/%s/trigger-attribution/%s/%s",
				r.Host,
				wellKnown,
				li[0],
				li[1])

			log.Println("Success execute conversion trigger(value 2)")
			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s, priority=%s", li[0], li[1])
	}

	if len(li) == 1 {
		if validation.IsValidTriggerData(li[0]) {
			location = fmt.Sprintf("http://%s/%s/trigger-attribution/%s",
				r.Host,
				wellKnown,
				li[0])

			log.Println("Success execute conversion trigger(value 1)")
			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s", li[0])
	}
	w.WriteHeader(http.StatusBadRequest)
}

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	beacon, err := ioutil.ReadFile(beaconPath)
	if err != nil {
		log.Printf("Ivalid beacon file path: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Success return beacon")
	w.Header().Set("Content-Type", "image/gif")
	w.Write(beacon)
	w.WriteHeader(http.StatusOK)
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
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

func PublicTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token, err := public_token.Generate()
	if err != nil {
		log.Printf("Failed generate public token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := database.InsertPublicToken(token, r.Referer(), r.Host); err != nil {
		log.Printf("Failed insert public token to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("{\"token_public_key\": \"%s\"}", token)

	log.Println("Success return public token")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
	w.WriteHeader(http.StatusOK)
}

func BlindSignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Invalid request method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var source *model.Source
	if err := json.NewDecoder(r.Body).Decode(source); err != nil {
		log.Printf("Invalid request body format: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, err := base64.RawURLEncoding.DecodeString(source.SourceToken)
	if err != nil {
		log.Printf("Failed base64URL decode source token: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signature, err := blind_sign.Execute(msg)
	if err != nil {
		log.Printf("Failed blind sign operation: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := database.InsertUnlinkableToken(source, signature, r.Referer(), r.Host); err != nil {
		log.Printf("Failed insert report to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("{\"unlinkable_token\": \"%s\"}", signature)

	log.Println("Success return unlinkable token")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
	w.WriteHeader(http.StatusOK)
}
