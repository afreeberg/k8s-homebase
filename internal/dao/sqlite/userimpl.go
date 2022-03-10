package sqlite

import "github.com/afreeberg/k8s-homebase/internal/models"

type UserImplSqlite struct {
}

func (dao UserImplSqlite) AutoMigrate() error {
	db := get()
	return db.AutoMigrate(&models.User{})
}

func (dao UserImplSqlite) Create(u *models.User) error {
	db := get()
	result := db.Create(u)
	return result.Error
}

func (dao UserImplSqlite) GetAll() ([]models.User, error) {
	db := get()
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}
