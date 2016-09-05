package main

import (
	"net/http"
	"encoding/json"
)

func Authentication(w http.ResponseWriter, r *http.Request) {

	user := User{
		Email: "vanessaschissato@gmail.com",
		Name:  "Vanessa Schissato",
		Phone: "+55 (11) 9 4187-8106",
	}

	json.NewEncoder(w).Encode(user)
}
