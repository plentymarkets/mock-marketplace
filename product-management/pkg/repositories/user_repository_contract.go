package repositories

import (
	"product-management/pkg/models"
)

type UserRepositoryContract interface {
	FetchByUser(user models.User) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
}
