package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB gorm.DB

type DatabaseConfig struct {
	User   string `json:"user"`
	DBname string `json:"dbname"`
}

func init() {
	file, err := os.Open("configs/database.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	config := &DatabaseConfig{}

	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}

	dbStr := fmt.Sprintf("user=%v dbname=%v sslmode=disable", config.User, config.DBname)

	DB, err = gorm.Open("postgres", dbStr)
	if err != nil {
		panic(err)
	}

	DB.DB()
	DB.DB().Ping()

	DB.AutoMigrate(&User{})
}
