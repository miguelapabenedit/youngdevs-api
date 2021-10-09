package data

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name              string
	Level             uint
	Map               string
	NumberOfColumns   uint
	NumberOfRows      uint
	AvailableCommands uint
	IsPremium         bool
	BestSolution      int
	BestTime          int
}
