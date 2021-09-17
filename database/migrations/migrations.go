package migrations

import (
	"gorm.io/gorm"

	"github.com/wallacemachado/api-bank-transfers/models"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Account{})
}
