package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"github/miguelapabenedit/youngdevs-api/app/data"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	server          = os.Getenv("DB_SERVER")
	port            = os.Getenv("DB_PORT")
	password        = os.Getenv("DB_PASS")
	user            = os.Getenv("DB_USER")
	database        = os.Getenv("DB_NAME")
	retryAttemps    = 0
	maxRetryAttemps = 3
)

func NewPostgreSQL() *gorm.DB {
	fmt.Println("Starting connection")
	db = openConnection()

	fmt.Println("Start DB Migration...")
	migrate()
	fmt.Println("Start DB Seed...")
	seed()

	return db
}

func openConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", server, user, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		retryConnection()
	}
	fmt.Println("Connection Established")
	return db
}

func migrate() {
	db.AutoMigrate(&data.User{})
	db.AutoMigrate(&data.Level{})
	db.AutoMigrate(&data.UserLevelState{})
}

func seed() {
	lvl1 := data.Level{
		Name:              "First Level",
		Level:             1,
		NumberOfColumns:   5,
		NumberOfRows:      5,
		AvailableCommands: "[0,1,2,3]",
		IsPremium:         false,
		Map:               "[[1,0,2,0,0],[0,0,2,0,0],[0,0,0,3,0],[0,0,2,0,0],[0,0,2,0,0]]",
	}
	lvl2 := data.Level{
		Name:              "Second Level",
		Level:             2,
		NumberOfColumns:   5,
		NumberOfRows:      5,
		AvailableCommands: "[0,1,2,3,4,6,7]",
		IsPremium:         false,
		Map:               "[[1,0,2,0,0],[0,0,0,0,0],[2,0,0,3,0],[0,0,2,0,0],[0,0,0,0,0]]",
	}
	lvl3 := data.Level{
		Name:              "Tirth Level",
		Level:             3,
		NumberOfColumns:   5,
		NumberOfRows:      5,
		AvailableCommands: "[2]",
		IsPremium:         true,
		Map:               "[[1,0,0,0,0],[0,0,0,0,0],[0,0,0,3,0],[0,0,0,0,0],[0,0,0,0,0]]",
	}

	db.Create(&lvl1)
	db.Create(&lvl2)
	db.Create(&lvl3)
}

func retryConnection() {
	if retryAttemps >= maxRetryAttemps {
		panic("Retry connection attempts exceded the max stablished")
	}

	time.AfterFunc(5*time.Second, func() {
		retryAttemps++
		log.Printf("Retrying to connect to the database %v/%v", retryAttemps, maxRetryAttemps)
		openConnection()
	})

}
