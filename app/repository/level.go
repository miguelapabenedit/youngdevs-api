package repository

import "github/miguelapabenedit/youngdevs-api/app/data"

type GetLevel interface {
	GetLevel(level int) *data.Level
}
