package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	var err error
	maxBytes := 1048576

	reader := io.LimitReader(r.Body, int64(maxBytes))

	dec := json.NewDecoder(reader)
	err = dec.Decode(data)
	if err != nil {
		return err
	}

	// Validate if the payload contains only a single json value or not
	// Check for extra data using peek
	if dec.More() {
		return errors.New("request body must contain only a single JSON value")
	}

	return nil
}

func (app *Config) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}
