package storage

import (
	"context"

	"calc_service/internal/models"
)

type Repository interface {
	SaveCalculation(ctx context.Context, calc models.Calculation) error
}
