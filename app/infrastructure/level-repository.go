package infrastructure

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
)

type levelRepo struct{}

func NewLevelRepository() repository.GetLevel {
	return &levelRepo{}
}

func (r *levelRepo) GetLevel(id int) *data.Level {
	var level data.Level

	result := db.Where("level", id).First(&level)

	if result.Error != nil {
		fmt.Println("An error has ocurred")
	}

	return &level
}