package interfaces

import "tma/services/account/internal/models"

type Repository interface {
	CreateUser(user models.User) (int, error)
	GetUserByLogin(login string) (models.User, error)
}
