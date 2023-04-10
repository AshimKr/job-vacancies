package initializers

import "example/challenge/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Job{}, &models.AllJob{})
}
