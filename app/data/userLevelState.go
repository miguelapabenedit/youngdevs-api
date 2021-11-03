package data

type UserLevelState struct {
	UserID       uint `gorm:"primaryKey;autoIncrement:false"`
	LevelID      uint `gorm:"primaryKey;autoIncrement:false"`
	Time         int
	UserSolution string
	IsSolved     bool
	Score        int
}
