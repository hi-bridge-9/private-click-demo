package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func topPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, wellKnown+"/trigger-attribution")
	li := filepath.SplitList(param)
	fmt.Println(li)
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func publicTokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func blindSignHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
