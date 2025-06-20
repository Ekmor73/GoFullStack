package database

import (
	"context"
	"go-fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Msg("Не удалось подключиться к базе данных")
		panic(err)
	}
	logger.Info().Msg("Подключились к базе данных")
	//logger.Info().Msgf("Подключаемся к БД: %s", config.Url)
	return dbpool
}
