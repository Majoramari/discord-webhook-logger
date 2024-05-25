package utils

import (
	"encoding/json"
	"net/http"

	"github.com/majoramari/visitor-logger/utils/webhook"
	"github.com/majoramari/visitor-logger/visitor"
)

func LogVisitorInfo(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	visitorInfo := visitor.ParseVisitorInfo(r, data)
	webhook.SendToWebhook(visitorInfo)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Visitor information logged successfully"))
}
