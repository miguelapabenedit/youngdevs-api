package repository

import "github/miguelapabenedit/youngdevs-api/app/data"

type GetUser interface {
	Get(authProviderId string) *data.User
}

type CreateUser interface {
	Create(u *data.User) error
}

type User interface {
	CreateUser
	GetUser
}
