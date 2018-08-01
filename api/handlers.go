package api

import (
	"encoding/json"
	"fmt"
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

	writeJSON(w, newCard)
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

	balance, blocked, err := cardService.Storage.GetBalance(card.CardID)
	if err != nil {
		log.Println("error getting balance:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{'error':'%s'}", err.Error())))
		return
	}

	writeJSON(w, struct {
		Balance float64
		Blocked float64
	}{balance, blocked})
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

	balance, err := cardService.Storage.Deposit(deposit.CardID, deposit.Amount)
	if err != nil {
		log.Println("error calling deposit function:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{'error':'%s'}", err.Error())))
		return
	}
	writeJSON(w, struct{ Total float64 }{balance})
}

// Merchants Handlers

// captureAuthHandler
// payload to process: {"card_id":1, "amount":1000}
func blockAuthHandler(w http.ResponseWriter, r *http.Request) {
	blocking := struct {
		CardID uint64  `json:"card_id"`
		Amount float64 `json:"amount"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&blocking); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{'error':'%s'}", err.Error())))
		return
	}

	blockingID, err := cardService.Storage.BlockAuthRequest(blocking.CardID, blocking.Amount)
	if err != nil {
		log.Println("error calling BlockAuthRequest:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{'error':'%s'}", err.Error())))
		return
	}

	writeJSON(w, struct{ BlockingID uint64 }{blockingID})
}

// captureHandler
// payload to process: {"card_id":1, "blocking_auth_id":1}
func captureHandler(w http.ResponseWriter, r *http.Request) {
	capture := struct {
		CardID     uint64 `json:"card_id"`
		BlockingID uint64 `json:"blocking_auth_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&capture); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	capturedAmount, err := cardService.Storage.CaptureRequest(capture.CardID, capture.BlockingID)
	if err != nil {
		log.Println("error calling CaptureRequest:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{'error':'%s'}", err.Error())))
		return
	}

	writeJSON(w, struct{ Captured float64 }{capturedAmount})
}

// cancelCaptureAuthHandler
// payload to process: {"card_id":1, "auth_number":1}
func cancelBlockingAuthHandler(w http.ResponseWriter, r *http.Request) {
	blockingAuth := struct {
		CardID     uint64 `json:"card_id"`
		BlockingID uint64 `json:"blocking_auth_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&blockingAuth); err != nil {
		log.Println("error decoding JSON payload:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := cardService.Storage.CancelCaptureAuth(blockingAuth.CardID, blockingAuth.BlockingID); err != nil {
		log.Println("error calling cancel capture auth", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("{'error':'%s'}", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	writeJSON(w, struct{ Result string }{"ok"})
}
