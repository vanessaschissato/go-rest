package main

import (
	"encoding/json"
	"net/http"
)

func Authentication(w http.ResponseWriter, r *http.Request) {

	user := User{
		Email: "vanessaschissato@gmail.com",
		Name:  "Vanessa Schissato",
		Phone: "+55 (11) 9 4321-1234",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}
