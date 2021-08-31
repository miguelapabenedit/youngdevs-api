package service

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/entity"
	"github/miguelapabenedit/youngdevs-api/app/repository"
)

type service struct{}

type Service interface {
	CreateUser(u *entity.User) error
}

var repo repository.Repository

func NewServices(repository repository.Repository) Service {
	repo = repository
	return &service{}
}

func (s *service) CreateUser(u *entity.User) error {
	err := repo.CreateUser(u)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
