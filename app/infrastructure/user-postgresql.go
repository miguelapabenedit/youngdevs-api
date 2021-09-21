package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct{}

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

func NewPostgreSQL() repository.User {
	fmt.Println("Starting connection")
	db = openConnection()
	fmt.Println("Start DB Migration...")
	db.AutoMigrate(&data.User{})
	return &repo{}
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

func (r *repo) CreateUser(u *data.User) error {
	return db.Create(u).Error
}

func (r *repo) GetUser(id string) *data.User {
	var user data.User

	result := db.First(&user)

	if result.Error != nil {
		fmt.Println("An error has ocurred")
	}

	return &user
}
