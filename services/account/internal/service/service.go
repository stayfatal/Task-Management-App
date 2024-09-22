package service

import (
	"context"
	"tma/services/account/internal/auth"
	"tma/services/account/internal/interfaces"
	"tma/services/account/internal/models"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo interfaces.Repository
}

func New(repo interfaces.Repository) interfaces.Service {
	return &service{repo: repo}
}

func (svc service) CreateAccount(ctx context.Context, user models.User) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "generating from password")
	}
	user.Password = string(hashPass)

	id, err := svc.repo.CreateUser(user)
	if err != nil {
		return "", errors.Wrap(err, "creating account repository level")
	}

	token, err := auth.CreateToken(id)
	if err != nil {
		return "", errors.Wrap(err, "creating token")
	}

	return token, nil
}

func (svc service) Login(ctx context.Context, user models.User) (string, error) {
	userFromRepo, err := svc.repo.GetUserByLogin(user.Login)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromRepo.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token, err := auth.CreateToken(userFromRepo.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
