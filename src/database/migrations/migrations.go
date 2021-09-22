package migrations

import (
	"gorm.io/gorm"

	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transfer{})

}

func DeleteTablesTestDb(db *gorm.DB) {
	db.Migrator().DropTable(&models.Account{})
	db.Migrator().DropTable(&models.Account{})

}
