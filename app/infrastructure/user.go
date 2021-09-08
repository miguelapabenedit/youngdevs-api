package infrastructure

import "github/miguelapabenedit/youngdevs-api/app/data"

func CreateUser(u *data.User) error {
	(u).ID = "idgenerated"
	return nil
}

func GetUser(id string) *data.User {
	var user data.User
	user.ID = id

	return &user
}
