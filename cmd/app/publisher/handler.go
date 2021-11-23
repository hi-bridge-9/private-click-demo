package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func topPageHandler(w http.ResponseWriter, r *http.Request) {

	// レスポンス
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, wellKnown+"/trigger-attribution")
	li := filepath.SplitList(param)
	fmt.Println(li)

	// レスポンス
	w.Header().Add("Location", "https://.....")
	w.WriteHeader(http.StatusFound)
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
