package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action       string          `json:"action"`
	Auth         AuthPld         `json:"auth,omitempty"`
	Registration RegistrationPld `json:"registration,omitempty"`
	Login        LoginPld        `json:"login,omitempty"`
}

type AuthPld struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationPld struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginPld struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit The Broker",
		Data:    nil,
	}
	_ = app.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPld RequestPayload
	err := app.ReadJSON(w, r, &requestPld)
	if err != nil {
		app.ErrorJSON(w, err)
	}

	switch requestPld.Action {
	case "auth":
		app.Authenticate(w, requestPld.Auth)
	case "registration":
		app.Register(w, requestPld.Registration)
	case "login":
		app.Login(w, requestPld.Login)

	default:
		app.ErrorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) Authenticate(w http.ResponseWriter, pld AuthPld) {
	// Create Some JSON and send it to auth service
	jsonData, _ := json.MarshalIndent(pld, "", "\t")
	log.Println("Making request To auth service: ")
	// Call the auth service
	request, err := http.NewRequest("POST", "http://authentication-service:8081/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()

	//Make sure we get back correct status code
	log.Println("Got response ", response.Body)
	if response.StatusCode == http.StatusUnauthorized {
		app.ErrorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.ErrorJSON(w, errors.New("error calling auth service"))
		return
	}

	// response body processing

	var jsonFromAuth JsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromAuth)
	if err != nil {
		app.ErrorJSON(w, errors.New("invalid credentials"))
		return
	}

	if jsonFromAuth.Error {
		app.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload JsonResponse

	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromAuth.Data

	app.WriteJSON(w, http.StatusAccepted, payload)
	return
}

func (app *Config) Register(w http.ResponseWriter, pld RegistrationPld) {
	// Create Some JSON and send it to auth service
	jsonData, _ := json.MarshalIndent(pld, "", "\t")
	log.Println("Making request To auth service: ")
	// Call the auth service
	request, err := http.NewRequest("POST", "http://authentication-service:8081/register", bytes.NewBuffer(jsonData))
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()

	//Make sure we get back correct status code
	log.Printf("Got response %+v", response.Body)
	if response.StatusCode == http.StatusUnauthorized {
		app.ErrorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.ErrorJSON(w, errors.New("error calling auth service"))
		return
	}

	// response body processing

	var jsonFromAuth JsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromAuth)
	if err != nil {
		app.ErrorJSON(w, errors.New("invalid credentials"))
		return
	}

	if jsonFromAuth.Error {
		app.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload JsonResponse

	payload.Error = false
	payload.Message = "Registration Complete"
	payload.Data = jsonFromAuth.Data

	app.WriteJSON(w, http.StatusCreated, payload)
	return
}

func (app *Config) Login(w http.ResponseWriter, pld LoginPld) {
	// Create Some JSON and send it to auth service
	jsonData, _ := json.MarshalIndent(pld, "", "\t")
	log.Println("Making request To auth service: ")
	// Call the auth service
	request, err := http.NewRequest("POST", "http://authentication-service:8081/login", bytes.NewBuffer(jsonData))
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()

	//Make sure we get back correct status code
	log.Printf("Got response %+v", response.Body)
	if response.StatusCode == http.StatusUnauthorized {
		app.ErrorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.ErrorJSON(w, errors.New("error calling auth service"))
		return
	}

	// response body processing

	var jsonFromAuth JsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromAuth)
	if err != nil {
		app.ErrorJSON(w, errors.New("invalid credentials"))
		return
	}

	if jsonFromAuth.Error {
		app.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload JsonResponse

	payload.Error = false
	payload.Message = "Login Complete"
	payload.Data = jsonFromAuth.Data

	app.WriteJSON(w, http.StatusCreated, payload)
	return
}
