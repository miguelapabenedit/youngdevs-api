package infrastructure

import "github/miguelapabenedit/youngdevs-api/app/data"

func CreateUser(u *data.User) error {
	(u).ID = "IDGeneratedYoungdevsApi:" + u.ID
	return nil
}

func GetUser(id string) *data.User {
	var user data.User
	user.ID = "IDGeneratedYoungdevsApi:" + id

	return &user
}
