package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kyu-takahahsi/private-click-demo/cmd/app/deliver/ad"
	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/validation"
)

func Delivery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	adTag, err := ad.Generate(r.UserAgent())
	if err != nil {
		log.Printf("Failed generate ad tag: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("Callback({'ads': '%s'});", adTag)

	log.Println("Success return ad callback function")
	w.Write([]byte(resp))
	w.WriteHeader(http.StatusOK)
	return
}

func Click(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	location := r.URL.Query().Get("lp")
	if location == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("Success redirect lp page")
	w.Header().Add("Location", location)
	w.WriteHeader(http.StatusFound)
	return
}

func Conversion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	triggerData := r.URL.Query().Get("cv")
	priority := r.URL.Query().Get("priority")
	if triggerData == "" {
		log.Println("Not exist trigger-data in request parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var location string
	if priority != "" {
		if validation.IsValidTriggerDataAndPriority(triggerData, priority) {
			location = fmt.Sprintf("%s/trigger/%s/%s", r.Host, triggerData, priority)

			log.Println("Success redirect media page(value 2)")
			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s, priority=%s", triggerData, priority)
	} else {
		if validation.IsValidTriggerData(triggerData) {
			location = fmt.Sprintf("%s/trigger/%s", r.Host, triggerData)

			log.Println("Success redirect media page(value 1)")
			w.Header().Add("Location", location)
			w.WriteHeader(http.StatusFound)
			return
		}
		log.Printf("Ivalid parameter value: trigger-data=%s", triggerData)
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
