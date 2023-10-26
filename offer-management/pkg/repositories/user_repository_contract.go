package repositories

import (
	"offer-management/pkg/models"
)

type UserRepositoryContract interface {
	FetchByID(id string) (models.User, error)
	FetchAll(page int, UsersPerPage int) ([]models.User, int, error)
	Create(offer models.User) (models.User, error)
	Update(offer models.User) (models.User, error)
}
