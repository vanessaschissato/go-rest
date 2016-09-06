package main

var currentId int

var users []User

// Give us some seed data
func init() {

	users = append(users, User{
		ID: "vanessaschissato@gmail.com",
		Email: "vanessaschissato@gmail.com",
		Name:  "Vanessa Schissato",
		Phone: "+55 (11) 9 4321-1234",
		password: "quilmes123",
	})
}

func RepoAuthenticateUser(credential Credential) User {

	for _, user := range users {
		if (user.Email == credential.Email) && (user.password == credential.Password) {
			return user
		}
	}
	// return empty User if not found
	return User{}
}