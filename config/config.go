package config

import (
	comment "ALTA_BE_SOSMED/features/comment/data"
	post "ALTA_BE_SOSMED/features/post/data"
	user "ALTA_BE_SOSMED/features/user/data"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var JWTSECRET = ""

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

func AssignEnv(c AppConfig) (AppConfig, bool) {
	var missing = false

	if val, found := os.LookupEnv("DBUSER"); found {
		c.DBUsername = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		c.DBPassword = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		c.DBPort = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		c.DBHost = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		c.DBName = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("JWT_SECRET"); found {
		JWTSECRET = val
	} else {
		missing = true
	}
	return c, missing
}

func InitConfig() AppConfig {
	var result AppConfig
	var missing = false
	result, missing = AssignEnv(result)
	if missing {
		godotenv.Load(".env")
		result, _ = AssignEnv(result)
	}

	return result
}

func InitSQL(c AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("terjadi error", err.Error())
		return nil
	}

	db.AutoMigrate(&user.User{}, &comment.Comment{}, &post.Post{})

	return db
}
