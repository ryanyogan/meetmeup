package postgres

import (
	"github.com/go-pg/pg/v9"
	"yogan.dev/meetmeup/pkg/models"
)

// UserRepo holds a db references
type UserRepo struct {
	DB *pg.DB
}

// GetUserByID takes an id and returns a user or an error
func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User

	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
