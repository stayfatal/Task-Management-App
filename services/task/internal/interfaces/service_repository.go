package interfaces

import "tma/services/task/internal/models"

type Service interface {
	AddTask(task models.Task) error
}
