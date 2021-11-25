package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/database"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/validation"
)

var (
	topPagePath = os.Getenv("TOP_PAGE_FILE_PATH")
	beaconPath  = os.Getenv("BEACON_FILE_PATH")
	wellKnown   = "/.well-known/private-click-measurement"
)

func TopPageHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンス
	t, err := template.ParseFiles(topPagePath)
	if err != nil {
		log.Printf("Ivalid template file path: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Printf("Failed template execute: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Success return html template")
}

func TriggerHandler(w http.ResponseWriter, r *http.Request) {
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

			log.Println("Success execute conversion trigger(2)")
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

			log.Println("Success execute conversion trigger(1)")
			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s", li[0])
	}
	w.WriteHeader(http.StatusNotFound)
}

func BeaconHandler(w http.ResponseWriter, r *http.Request) {
	beacon, err := ioutil.ReadFile(beaconPath)
	if err != nil {
		log.Printf("Ivalid beacon file path: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// レスポンス
	log.Println("Success return beacon")
	w.Header().Set("Content-Type", "image/gif")
	w.Write(beacon)
	w.WriteHeader(http.StatusOK)
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		log.Printf("Failed connect to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	insert := database.GenerateInsertReportQuery(r)
	_, err = db.Exec(insert)
	if err != nil {
		log.Printf("Failed report data insert to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// レスポンス
	log.Println("Success report data insert to DB")
	w.WriteHeader(http.StatusOK)
}

func PublicTokenHandler(w http.ResponseWriter, r *http.Request) {
	// 公開鍵を取得、BASE64URLエンコード、JSON形式の文字列を出力
	var publicToken string

	db, err := database.Connect()
	if err != nil {
		log.Printf("Failed connect to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	insert := database.GenerateInsertPublicTokenQuery(publicToken, r)
	_, err = db.Exec(insert)
	if err != nil {
		log.Printf("Failed public token insert to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Success public token insert to DB")

	// レスポンス
	log.Println("Success return public token")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(publicToken))
	w.WriteHeader(http.StatusOK)
}

func BlindSignHandler(w http.ResponseWriter, r *http.Request) {
	// 署名を出力、BASE64URLエンコード、JSON形式の文字列を出力
	var signature string

	db, err := database.Connect()
	if err != nil {
		log.Printf("Failed connect to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	insert := database.GenerateInsertUnlinkableTokenQuery(signature, r)
	_, err = db.Exec(insert)
	if err != nil {
		log.Printf("Failed unlinkable token insert to DB: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Success unlinkable token insert to DB")

	// レスポンス
	log.Println("Success return unlinkable token")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(signature))
	w.WriteHeader(http.StatusOK)
}
