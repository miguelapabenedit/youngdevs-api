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
}

type UsersPaginated struct {
	TotalCount int64
	PageSize   int
	PageIndex  int
	Users      []User
}
