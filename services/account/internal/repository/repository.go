package repository

import (
	"tma/services/account/internal/interfaces"
	"tma/services/account/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.Repository {
	return &repository{db: db}
}

func (repo *repository) CreateUser(user models.User) (int, error) {
	var id int
	err := repo.db.QueryRow("insert into users (login,password) values ($1,$2) returning id", user.Login, user.Password).Scan(&id)
	return id, errors.Wrap(err, "inserting into db")
}

func (repo *repository) GetUserByLogin(login string) (models.User, error) {
	user := models.User{}
	err := repo.db.QueryRow("select * from users where login = $1", login).Scan(&user.Id, &user.Login, &user.Password)
	return user, errors.Wrap(err, "selecting from db")
}
