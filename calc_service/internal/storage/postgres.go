package storage

import (
	"context"
	"database/sql"

	"calc_service/internal/models"

	"github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p *PostgresRepository) SaveCalculation(
	ctx context.Context,
	calc models.Calculation,
) error {

	query := `
		INSERT INTO calculations(
			operation,
			operands,
			result,
			request_id,
			user_id,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := p.db.ExecContext(
		ctx,
		query,
		calc.Operation,
		pq.Array(calc.Operands),
		calc.Result,
		calc.RequestID,
		calc.UserID,
		calc.CreatedAt,
	)

	return err

}
