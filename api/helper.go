package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Helper function to write json file in response writer
func writeJSON(w http.ResponseWriter, data interface{}) {
	jData, err := json.Marshal(data)
	if err != nil {
		log.Println("error serializing data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
