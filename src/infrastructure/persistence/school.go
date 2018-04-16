package persistence

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
)

type schoolRepository struct {
	DBClient *sqlx.DB
}

func NewSchoolRepository(dbc *sqlx.DB) repository.SchoolRepository {
	return &schoolRepository{
		DBClient: dbc,
	}
}

func (repo *schoolRepository) GetSchools(ctx context.Context) ([]*entity.School, error) {
	sql, _, _ := sq.Select("*").From("schools").ToSql()
	schools := []entity.School{}
	err := repo.DBClient.SelectContext(ctx, &schools, sql)
	if err != nil {
		panic(err)
	}

	_schools := []*entity.School{}
	for i := range schools {
		_schools = append(_schools, &schools[i])
	}
	return _schools, nil
}
