package initializers

import "go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Movie{})
	DB.AutoMigrate(&models.NonMovie{})
}
