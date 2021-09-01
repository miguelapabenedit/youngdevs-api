package repository

import "github/miguelapabenedit/youngdevs-api/app/entity"

type repository struct{}

type Repository interface {
	CreateUser(u *entity.User) error
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateUser(u *entity.User) error {
	(u).ID = "idgenerated"
	return nil

}

func GetUser(id string) entity.User {
	var user entity.User
	user.ID = id

	return user
}
