package persistence

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
)

type classRepository struct {
	DBClient *sqlx.DB
}

func NewClassRepository(dbc *sqlx.DB) repository.ClassRepository {
	return &classRepository{
		DBClient: dbc,
	}
}

func (repo *classRepository) GetClasses(ctx context.Context, schoolID int64) ([]*entity.Class, error) {
	sql, args, _ := sq.Select("*").From("classes").Where(sq.Eq{"school_id": schoolID}).ToSql()
	classes := []entity.Class{}
	err := repo.DBClient.SelectContext(ctx, &classes, sql, args...)
	if err != nil {
		panic(err)
	}

	_classes := []*entity.Class{}
	for i := range classes {
		_classes = append(_classes, &classes[i])
	}
	return _classes, nil
}
