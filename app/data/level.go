package data

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	Name              string
	Level             uint
	Map               string
	NumberOfColumns   uint
	NumberOfRows      uint
	AvailableCommands string
	IsPremium         bool
	bestSolution      int
	bestTime          int
}
