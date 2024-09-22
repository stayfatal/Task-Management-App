package interfaces

import (
	"context"
	"tma/services/account/internal/models"
)

type Service interface {
	CreateAccount(ctx context.Context, user models.User) (string, error)
	Login(ctx context.Context, user models.User) (string, error)
}
