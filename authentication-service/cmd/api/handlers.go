package main

import (
	"authentication/data"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

const (
	ErrRecordNotFound = "sql: no rows in result set"
)

type (
	authPld struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	regPld struct {
		Email     string `json:"email"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Password  string `json:"password"`
	}

	loginPld struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload authPld
	log.Println("Got Hit from broker")
	err := app.ReadJSON(w, r, &requestPayload)
	log.Println("Received payload ", requestPayload)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	//	validate the user
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}
	log.Println("found user :", user)
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		log.Println("Password did not match")
		app.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
}

func (app *Config) Register(w http.ResponseWriter, r *http.Request) {
	var requestPayload regPld
	log.Println("Got Hit from broker")
	err := app.ReadJSON(w, r, &requestPayload)
	log.Println("Received payload ", requestPayload)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	//	validate the user
	user, err := app.Models.User.GetByEmail(requestPayload.Email)

	if err != nil && err.Error() != ErrRecordNotFound {
		app.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	if user != nil {
		log.Println("found user :", user)
		app.ErrorJSON(w, errors.New("user already exist"), http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestPayload.Password), bcrypt.DefaultCost)
	u := data.User{
		Email:     requestPayload.Email,
		FirstName: requestPayload.FirstName,
		LastName:  requestPayload.LastName,
		Password:  string(hashedPassword),
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.Create()
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Println("User created successfully")

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created user %s", u.Email),
		Data:    u,
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
}

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	var body loginPld
	log.Println("Got Hit from broker")
	err := app.ReadJSON(w, r, &body)
	log.Println("Received payload ", body)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	//	validate the user
	user, err := app.Models.User.GetByEmail(body.Email)

	if err != nil && err.Error() != ErrRecordNotFound {
		app.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	log.Println("found user :", user)
	valid, err := user.PasswordMatches(body.Password)
	if err != nil || !valid {
		log.Println("Password did not match")
		app.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}
	log.Println("Successfully retrieved User and Logged into system.")
	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
}
