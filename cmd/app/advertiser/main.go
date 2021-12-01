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
	// リクエスト時のパスとハンドラーのマッピング
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		// 広告主のLPを返却
		"/lp": handler.LandingPageHandler,

		// 広告主のCV地点ページを返却
		"/cv/": handler.CVPageHandler,
	}

	log.Println("Web serever start")
	wb := server.NewWebServer(handlerMap)
	if err := wb.Start(port); err != nil {
		log.Fatal(err)
	}
}
