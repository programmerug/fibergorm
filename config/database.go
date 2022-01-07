// 文件名：config/database.go
package config

import (
	"github.com/programmerug/fibergorm/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	Database, err = gorm.Open(sqlite.Open("fibergorm.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Dog{})

	return nil
}
