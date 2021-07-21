package pg

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"mayerkv/otus-ma-lesson-6/domain"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `insert into users (id, username, first_name, last_name, email, phone) 
values ($1, $2, $3, $4, $5, $6) 
on conflict (id) do update
set username = excluded.username,
    first_name = excluded.first_name,
    last_name = excluded.last_name,
    email = excluded.email,
    phone = excluded.phone`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.Phone)

	return handlerError(err)
}

func (r *UserRepository) FindById(id int) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user domain.User
	query := `select id, username, first_name, last_name, email, phone from users where id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &user, nil
	default:
		return nil, handlerError(err)
	}
}

func (r *UserRepository) Delete(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `delete from users where id = $1`
	_, err := r.db.ExecContext(ctx, query, user.ID)

	return handlerError(err)
}

func (r *UserRepository) NextId() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var id int
	query := "select nextval('users_id_seq')"
	err := r.db.GetContext(ctx, &id, query)
	if err != nil {
		return 0, handlerError(err)
	}

	return id, nil
}
