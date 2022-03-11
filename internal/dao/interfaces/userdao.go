package interfaces

import "github.com/afreeberg/k8s-homebase/internal/models"

type UserDao interface {
	AutoMigrate() error
	Create(u *models.User) error
	Update(u *models.User) error
	Delete(u *models.User) error
	GetById(id int) (models.User, error)
	GetAll() ([]models.User, error)
}
