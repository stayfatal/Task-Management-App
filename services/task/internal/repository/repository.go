package repository

import (
	"tma/services/task/internal/interfaces"
	"tma/services/task/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.Repository {
	return &repository{db: db}
}

func (repo *repository) CreateTask(task models.Task) error {
	_, err := repo.db.Exec("insert into tasks (name,description,created_at,deadline) values ($1,$2,$3,$4)", task.Name, task.Description, task.CreatedAt, task.Deadline)
	return errors.Wrap(err, "inserting into db")
}
