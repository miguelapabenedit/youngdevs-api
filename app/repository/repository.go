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
