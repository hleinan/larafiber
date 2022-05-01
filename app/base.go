package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

type GormLogger struct{}

var db *gorm.DB //database

type Base struct {
	ID        int64      `gorm:"primary_key" json:"-"`
	UUID      string     `gorm:"type:string;" json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	uuidGen := uuid.Must(uuid.NewRandom()).String()
	base.UUID = strings.ReplaceAll(uuidGen, "-", "")
	return nil
}

func init() {
	logFile, err := os.OpenFile("gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // Disable color
		},
	)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := gorm.Open(sqlite.Open(os.Getenv("db_path")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})

	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(
		&ForgottenPassword{},
		&User{},
		&Account{},
		&Version{},
		&Task{},
	)
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("What?")
	} else {
		fmt.Println(sqlDB.Stats().OpenConnections)
	}
	//defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	//defer db.
}
