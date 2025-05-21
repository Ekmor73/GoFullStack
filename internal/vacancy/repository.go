package vacancy

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancRepository {
	return &VacancRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *VacancRepository) addVacancy(form VacancRepository) {

}
