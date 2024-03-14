package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit The Broker",
		Data:    nil,
	}
	_ = app.WriteJSON(w, http.StatusOK, payload)
}
