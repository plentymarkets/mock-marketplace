package repositories

import (
	"offer-management/pkg/models"
)

type UserRepositoryContract interface {
	FetchAll(page int, UsersPerPage int) ([]models.User, int, error)
	FetchByID(id string) (models.User, error)
	FetchByName(UserName string) (models.User, error)
	Create(offer models.User) (models.User, error)
	Update(offer models.User) (models.User, error)
}
