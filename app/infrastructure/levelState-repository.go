package infrastructure

import (
	"errors"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
)

type userLevelStateRepo struct{}

func NewLevelStateRepository() repository.UserLevelState {
	return &userLevelStateRepo{}
}

func (r *userLevelStateRepo) CreateLevelState(u *data.UserLevelState) error {
	var uls data.UserLevelState

	result := db.Where("user_id = ? and level_id = ?", u.UserID, u.LevelID).First(&uls)

	if result.Error != nil && result.Error.Error() != "record not found" {
		fmt.Println("the levelstate is already register")
		return errors.New("an error has ocured")
	}

	if result.RowsAffected != 0 {
		fmt.Println("the levelstate is already register")
		return errors.New("the levelstate is already register")
	}

	return db.Create(&u).Error
}

func (r *userLevelStateRepo) GetAllUserLevelState(userId int) []data.UserLevelState {
	var uls []data.UserLevelState

	db.Where("user_id = ?", userId).Find(&uls)

	return uls
}

func (r *userLevelStateRepo) DeleteAllById(userId uint) error {
	result := db.Where("user_id = ?", userId).Delete(data.UserLevelState{})

	return result.Error
}

func (r *userLevelStateRepo) UpdateLevelState(u *data.UserLevelState) error {
	if !u.IsSolved {
		var uls data.UserLevelState

		db.Where("user_id = ? and level_id = ?", u.UserID, u.LevelID).First(&uls)

		u.IsSolved = uls.IsSolved
	}

	return db.Save(&u).Error
}

func (r *userLevelStateRepo) GetLevelState(u *data.UserLevelState) error {
	result := db.Where("user_id = ? and level_id = ?", u.UserID, u.LevelID).First(&u)

	if result.Error != nil {
		fmt.Println("An error has ocurred")
		return result.Error
	}

	return nil
}
