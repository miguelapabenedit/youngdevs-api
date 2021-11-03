package repository

import "github/miguelapabenedit/youngdevs-api/app/data"

type GetUserLevelState interface {
	GetLevelState(u *data.UserLevelState) error
}

type CreateUserLevelState interface {
	CreateLevelState(u *data.UserLevelState) error
}

type UpdateUserLevelState interface {
	UpdateLevelState(u *data.UserLevelState) error
}

type UserLevelState interface {
	GetUserLevelState
	CreateUserLevelState
	UpdateUserLevelState
}
