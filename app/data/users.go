package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AuthProviderUserId string
	Email              string
	IsPremium          bool
	IsLocked           bool
	IsAdmin            bool
	Score              int
	CurrentLevel       uint
}

type UsersPaginated struct {
	TotalCount int64
	PageSize   int
	PageIndex  int
	Users      []User
}
