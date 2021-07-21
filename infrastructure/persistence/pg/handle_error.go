package pg

import (
	"github.com/jackc/pgconn"
	"mayerkv/otus-ma-lesson-6/domain"
)

func handlerError(err error) error {
	if pqError, ok := err.(*pgconn.PgError); ok {
		return handleDriverError(pqError)
	}

	return err
}

func handleDriverError(err *pgconn.PgError) error {
	switch err.Code {
	case "23505":
		return domain.NewInvalidArgument(err.Message)
	default:
		return err
	}
}

