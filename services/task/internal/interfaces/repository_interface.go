package interfaces

import "tma/services/task/internal/models"

type Repository interface {
	CreateTask(task models.Task) error
}
