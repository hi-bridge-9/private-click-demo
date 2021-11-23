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
	// リクエスト時のパスとハンドラーのマッピング
	handlerMap := map[string]func(w http.ResponseWriter, r *http.Request){
		// 広告配信面の返却
		"/top-page/": topPageHandler,

		// CVトリガーリダイレクト実行（CVタグからのリクエスト）
		"/trigger/": triggerHandler,

		// CVトリガーリダイレクトに対して画像の返却
		wellKnown + "/trigger-attribution/": beaconHandler,

		// レポートの受け取り
		wellKnown + "/report-attribution/": reportHandler,

		// ---------（任意）アドフラウド防止機能---------
		// 公開鍵の返却
		wellKnown + "/get-token-public-key/": publicTokenHandler,

		// ブラインド署名の実行、署名の返却
		wellKnown + "/sign-unlinkable-token/": blindSignHandler,
		// -------------------------------------
	}

	// webサーバー起動
	ws := server.NewWebServer(handlerMap)
	if err := ws.Start(port); err != nil {
		log.Fatal(err)
	}
}
