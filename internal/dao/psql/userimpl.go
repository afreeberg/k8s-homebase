package psql

import "github.com/afreeberg/k8s-homebase/internal/models"

type UserImplPsql struct {
}

func (dao UserImplPsql) AutoMigrate() error {
	db := get()
	return db.AutoMigrate(&models.User{})
}

func (dao UserImplPsql) Create(u *models.User) error {
	db := get()
	result := db.Create(u)
	return result.Error
}

func (dao UserImplPsql) Update(u *models.User) error {
	db := get()
	result := db.Save(u)
	return result.Error
}

func (dao UserImplPsql) Delete(u *models.User) error {
	db := get()
	result := db.Delete(u)
	return result.Error
}

func (dao UserImplPsql) GetById(i int) (models.User, error) {
	db := get()
	var user models.User
	result := db.First(&user, i)
	return user, result.Error
}

func (dao UserImplPsql) GetAll() ([]models.User, error) {
	db := get()
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}
