package util

import (
	"Project/research/sample-gql/config"
	"Project/research/sample-gql/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDriver(config *config.AppConfig) *gorm.DB {

	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}

	DatabaseMigration(db)

	return db
}

func DatabaseMigration(db *gorm.DB) {
	// db.Migrator().DropTable(entities.Book{})
	// db.Migrator().DropTable(entities.Person{})
	db.AutoMigrate(entities.Person{})
	db.AutoMigrate(entities.Book{})

}
