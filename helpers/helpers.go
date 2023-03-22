package helpers

import (
	"encoding/json"
	"net/http"
)

func ClientError(w http.ResponseWriter, errorMsg string, status int) {
	data := map[string]string{"error": errorMsg}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

func ServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
