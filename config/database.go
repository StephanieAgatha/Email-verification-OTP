package config

import (
	"fmt"
	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"login-register-email/model"
)

var DB *gorm.DB

var (
	host     = ""
	user     = "postgres"
	password = ""
	port     = 6086
	dbname   = "railway"
)

func InitializeDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Warnf("Failed connect to database,reason :", err.Error())
		return
	}
	db.AutoMigrate(&model.Student{})

	slog.Infof("DB Connected to %v", host)
	DB = db
}
