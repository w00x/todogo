package gorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todogohexa/shared"
)

type PostgresBase struct {
	DB *gorm.DB
}

var dbConn *gorm.DB

func NewPostgresBase() *PostgresBase {
	connectionDb, err := getConnection()
	if err != nil {
		log.Panic(err)
	}
	return &PostgresBase{ connectionDb }
}

func getConnection() (*gorm.DB, error) {
	if dbConn == nil {
		fmt.Println("=========================NEW CONN DB=============================")
		uri := shared.GetEnv("DATABASE_URI")
		var err error
		dbConn, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	return dbConn, nil
}

