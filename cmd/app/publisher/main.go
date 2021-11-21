package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/server"
)

var (
	port = os.Getenv("PORT")

	// Private Click Measurementを利用するためのパス
	wellKnown = "/.well-known/private-click-measurement"
)

// SSP起動！！（トリガーオン的な）
func main() {
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		// 基本機能
		"/top-page":                         topPageHandler,
		wellKnown + "/trigger-attribution/": triggerHandler,
		wellKnown + "/report-attribution/":  reportHandler,

		// （任意）アドフラウド防止機能
		wellKnown + "/get-token-public-key/":  publicTokenHandler,
		wellKnown + "/sign-unlinkable-token/": blindSignHandler,
	}

	ws := server.NewWebServer(handlerMap)
	if err := ws.Start(port); err != nil {
		log.Fatal(err)
	}
}

func topPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
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


