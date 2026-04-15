package service

import "invoice_gen_be/internal/model"

var users = []model.User{
	{ID: 1, Username: "admin", Password: "admin123", Role: "Admin"},
	{ID: 2, Username: "kerani", Password: "kerani123", Role: "Staff"},
}

func Authenticate(username, password string) (*model.User, bool) {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return &u, true
		}
	}
	return nil, false
}


