package repo

import (
	"context"
	"market/internal/model"
)

type CoinRepo interface {
	FindByUnit(ctx context.Context, unit string) (coin *model.Coin, err error)
}
