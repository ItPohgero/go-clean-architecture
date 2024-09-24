package repository

import (
	"go-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HealthRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewHealthRepository(log *logrus.Logger) *HealthRepository {
	return &HealthRepository{
		Log: log,
	}
}

func (r *HealthRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}
