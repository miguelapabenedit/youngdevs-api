package infrastructure

import (
	"errors"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
)

type userRepo struct{}

func NewUserRepository() repository.User {
	return &userRepo{}
}

func (r *userRepo) Create(u *data.User) error {
	var user data.User

	result := db.Where("email = ?", u.Email).First(&user)

	if result.Error != nil && result.Error.Error() != "record not found" {
		fmt.Println("the email is already register")
		return errors.New("an error has ocured")
	}

	if user.ID != 0 {
		fmt.Println("the email is already register")
		return errors.New("the email is already register")
	}

	return db.Create(&u).Error
}

func (r *userRepo) Get(id string) *data.User {
	var user data.User

	result := db.Where("auth_provider_user_id = ?", id).First(&user)

	if result.Error != nil {
		fmt.Println("An error has ocurred")
	}

	return &user
}
