package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyu-takahahsi/private-click-demo/cmd/app/publisher/handler"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/server"
)

var (
	port = os.Getenv("WEB_SERVER_PORT")

	// Private Click Measurementを利用するためのパス
	wellKnown = "/.well-known/private-click-measurement"
)

func main() {
	// リクエスト時のパスとハンドラーのマッピング
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		// 広告配信面の返却
		"/top-page": handler.TopPage,

		// CVトリガーリダイレクト実行（CVタグからのリクエスト）
		"/cv-trigger/": handler.TriggerRedirect,

		// CVトリガーリダイレクトに対して画像の返却
		wellKnown + "/trigger-attribution/": handler.Beacon,

		// レポートの受け取り
		wellKnown + "/report-attribution/": handler.Report,

		// ---------（任意）アドフラウド防止機能---------
		// 公開鍵の返却
		wellKnown + "/get-token-public-key/": handler.PublicToken,

		// ブラインド署名の実行、署名の返却
		wellKnown + "/sign-unlinkable-token/": handler.BlindSign,
		// -------------------------------------
	}

	log.Println("Web serever start")
	ws := server.NewWebServer(handlerMap)
	if err := ws.Start(port); err != nil {
		log.Fatal(err)
	}
}
