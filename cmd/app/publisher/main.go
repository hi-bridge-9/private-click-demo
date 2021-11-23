package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/server"
)

var (
	port = os.Getenv("WEB_SERVER_PORT")

	// Private Click Measurementを利用するためのパス
	wellKnown = "/.well-known/private-click-measurement"
)

func main() {
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		// 広告配信面
		"/top-page":                         topPageHandler,

		// Private Click Measurement
		// 基本機能
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
