package repository

import "github/miguelapabenedit/youngdevs-api/app/data"

type GetUser interface {
	Get(authProviderId string) *data.User
}

type CreateUser interface {
	Create(u *data.User) error
}

type UpdateUser interface {
	Update(u *data.User) error
}

type GetAllUsers interface {
	GetAllWithPagination(users *data.UsersPaginated)
}

type User interface {
	CreateUser
	GetUser
	GetAllUsers
	UpdateUser
}
