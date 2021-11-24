package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/val"
)

func topPageHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンス
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/cv-trigger/")
	li := strings.Split(param, "/")

	if len(li) == 0 || len(li) > 2 {
		log.Printf("Ivalid parameter length: %d", len(li))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var location string
	if len(li) == 2 {
		if val.IsValidTriggerDataAndPriority(li[0], li[1]) {
			location = fmt.Sprintf("http://%s/%s/trigger-attribution/%s/%s",
				r.Host,
				wellKnown,
				li[0],
				li[1])

			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s, priority=%s", li[0], li[1])
	}

	if len(li) == 1 {
		if val.IsValidTriggerData(li[0]) {
			location = fmt.Sprintf("http://%s/%s/trigger-attribution/%s",
				r.Host,
				wellKnown,
				li[0])

			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s", li[0])
	}

	w.WriteHeader(http.StatusNotFound)
}

func beaconHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, wellKnown+"/trigger-attribution")
	li := filepath.SplitList(param)
	fmt.Println(li)

	// レスポンス
	w.Header().Set("Content-Type", "image/gif")
	w.Write([]byte("gif beacon"))
	w.WriteHeader(http.StatusOK)
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンス
	w.WriteHeader(http.StatusOK)
}

func publicTokenHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("public token"))
	w.WriteHeader(http.StatusOK)
}

func blindSignHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("unlinkable token"))
	w.WriteHeader(http.StatusOK)
}
