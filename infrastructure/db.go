package infrastructure

import (
	"fmt"

	"github.com/shankalpa12/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env *config.Config) Database {

	dsn := "sqlserver://gorm:shankalpa.koirala@localhost:3306?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ERROR WAS THROWN WHILE SETTING UP THE SERVER!!!!!!!")
	}

	return Database{DB: db}
}
