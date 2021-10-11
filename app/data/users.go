package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AuthProviderUserId string
	Email              string
	IsPremium          bool
	IsLocked           bool
	IsAdmin            bool
}
