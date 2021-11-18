package main

import (
	"fmt"
	"net/http"
)

// SSP起動！！（トリガーオン的な）
func main() {
	http.HandleFunc("/advertiser/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
