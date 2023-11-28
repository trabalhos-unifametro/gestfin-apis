package models

import (
	"gestfin-apis/database"
	"gestfin-apis/utils"
)

func RunMigrations() {
	db := database.OpenConnection()
	if err := db.AutoMigrate(User{}); err != nil {
		utils.LogMessage{Title: "[MIGRATIONS] Error on db.AutoMigrate(User{})", Body: err.Error()}.Error()
	}
	if err := db.AutoMigrate(Expenses{}); err != nil {
		utils.LogMessage{Title: "[MIGRATIONS] Error on db.AutoMigrate(Patient{})", Body: err.Error()}.Error()
	}
	if err := db.AutoMigrate(Goals{}); err != nil {
		utils.LogMessage{Title: "[MIGRATIONS] Error on db.AutoMigrate(Unit{})", Body: err.Error()}.Error()
	}
}
