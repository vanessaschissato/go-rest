package main

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"userID"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	password string
}
