package migrations

import (
	"github.com/Kovokar/BookApiGo/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}
