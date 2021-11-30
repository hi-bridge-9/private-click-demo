package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyu-takahahsi/private-click-demo/cmd/app/advertiser/handler"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/server"
)

var (
	port = os.Getenv("WEB_SERVER_PORT")
)

func main() {
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		"/lp":  handler.LandingPageHandler,
		"/cv/": handler.CVPageHandler,
	}

	wb := server.NewWebServer(handlerMap)
	if err := wb.Start(port); err != nil {
		log.Fatal(err)
	}
	log.Println("Web server start")
}
