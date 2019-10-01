package query

import (
	"database/sql"
	"structs"
)

func CheckValid(login, password string, database *sql.DB) bool {
	var users = []structs.Users{}
	users = QueryUsers(users, database)
	var valid bool = false
	for _, p := range users {
		if p.Login == login && p.Password == password {
			valid = true
			break
		}
	}
	return valid
}