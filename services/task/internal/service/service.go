package service

import (
	"time"
	"tma/services/task/internal/interfaces"
	"tma/services/task/internal/models"

	"github.com/pkg/errors"
)

type service struct {
	repo interfaces.Repository
}

func New(repo interfaces.Repository) interfaces.Service {
	return &service{repo: repo}
}

func (svc *service) AddTask(task models.Task) error {
	task.CreatedAt = time.Now()

	err := svc.repo.CreateTask(task)
	return errors.Wrap(err, "creating task repository level")
}
