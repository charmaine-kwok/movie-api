package initializers

import "go-crud/models"

func SyncDatabase() {
	CRED_DB.AutoMigrate(&models.User{})
}
