package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Clients Handlers
// newClientHandler - Method: POST
// payload to process: {"name":"christian espinoza"}
func newClientHandler(w http.ResponseWriter, r *http.Request) {
	client := struct {
		NameOnCard string `json:"name"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(client)

	newCard, err := cardService.Storage.New(client.NameOnCard)
	if err != nil {
		log.Println("error generating card")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(newCard)
	if err != nil {
		log.Println("error serializing card")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// balanceHandler - Method: Get
// payload to process: {"card_id":1}
func balanceHandler(w http.ResponseWriter, r *http.Request) {
	card := struct {
		CardID uint64 `json:"card_id"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(card)

	balance, err := cardService.Storage.GetBalance(card.CardID)
	if err != nil {
		log.Println("error getting balance:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(struct{ Balance float64 }{balance})
	if err != nil {
		log.Println("error serializing card")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// depositHandler - Method: POST
// payload to process: {"card_id":1, "amount":1000}
func depositHandler(w http.ResponseWriter, r *http.Request) {
	deposit := struct {
		CardID uint64  `json:"card_id"`
		Amount float64 `json:"amount"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(deposit)
	w.WriteHeader(http.StatusOK)
}

// Merchants Handlers

// captureAuthHandler
// payload to process: {"card_id":1, "amount":1000}
func captureAuthHandler(w http.ResponseWriter, r *http.Request) {
	deposit := struct {
		CardID uint64  `json:"card_id"`
		Amount float64 `json:"amount"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(deposit)
	w.WriteHeader(http.StatusOK)

}

// deleteCaptureAuthHandler
// payload to process: {"card_id":1, "auth_number":1}
func deleteCaptureAuthHandler(w http.ResponseWriter, r *http.Request) {
	deposit := struct {
		CardID     uint64  `json:"card_id"`
		AuthNumber float64 `json:"auth_number"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(deposit)
	w.WriteHeader(http.StatusOK)

}

// captureHandler
// payload to process: {"card_id":"1", "amount":1000}
func captureHandler(w http.ResponseWriter, r *http.Request) {
	deposit := struct {
		CardID uint64  `json:"card_id"`
		Amount float64 `json:"amount"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&deposit); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(deposit)
	w.WriteHeader(http.StatusOK)
}
