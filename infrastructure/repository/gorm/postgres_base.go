package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todogo/shared"
)

type PostgresBase struct {
	DB *gorm.DB
}

var DbConn *gorm.DB

func NewPostgresBase() *PostgresBase {
	connectionDb, err := getConnection()
	if err != nil {
		log.Panic(err)
	}
	return &PostgresBase{connectionDb}
}

func getConnection() (*gorm.DB, error) {
	if DbConn == nil {
		uri := shared.GetEnv("DATABASE_URI")
		var err error
		DbConn, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	return DbConn, nil
}
