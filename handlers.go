package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"
)

type APIError struct {
	Message       string `json:"message"`
}

func Authentication(w http.ResponseWriter, r *http.Request) {

	var credential Credential

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &credential); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			return
		}
	}

	if (credential.Email == "") {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{Message: "'email' is required"})
		return
	}

	if (credential.Password == "") {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{Message: "'password' is required"})
		return
	}

	authenticatedUser := RepoAuthenticateUser(credential)

	if (authenticatedUser == User{}) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(authenticatedUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
