package repository

import "github/miguelapabenedit/youngdevs-api/app/data"

type GetUser interface {
	GetUser(id string) *data.User
}

type CreateUser interface {
	CreateUser(u *data.User) error
}

type User interface {
	CreateUser
	GetUser
}
