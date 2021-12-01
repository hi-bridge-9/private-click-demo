package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyu-takahahsi/private-click-demo/cmd/app/deliver/handler"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/server"
)

var (
	port = os.Getenv("WEB_SERVER_PORT")
)

func main() {
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		// 広告情報の返却
		"/delivery": handler.Delivery,

		// 広告遷移先へのリダイレクト
		"/click": handler.Click,

		// 広告の配信面へのリダイレクト
		"/conversion": handler.Conversion,
	}

	log.Println("Web serever start")
	wb := server.NewWebServer(handlerMap)
	if err := wb.Start(port); err != nil {
		log.Fatal(err)
	}
}
