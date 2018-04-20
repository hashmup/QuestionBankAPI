package persistence

import (
	"context"
	"time"

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

func (repo *classRepository) GetClassesBySchoolID(ctx context.Context, schoolID int64) ([]*entity.Class, error) {
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

func (repo *classRepository) GetClassesByUserID(ctx context.Context, userID int64) ([]*entity.Class, error) {
	sql, args, _ := sq.Select("classes.*").From("user_class_assoc").Join("classes ON user_class_assoc.class_id = classes.id").Where(sq.Eq{"user_classes.user_id": userID}).ToSql()
	classes := []entity.Class{}
	err := repo.DBClient.SelectContext(ctx, &classes, sql, args...)
	if err != nil {
		return nil, err
	}

	_classes := []*entity.Class{}
	for i := range classes {
		_classes = append(_classes, &classes[i])
	}
	return _classes, nil
}

func (repo *classRepository) GetFolders(ctx context.Context, classID int64) ([]*entity.Folder, error) {
	sql, args, _ := sq.Select("*").From("folders").Where(sq.Eq{"class_id": classID}).ToSql()
	folders := []entity.Folder{}
	err := repo.DBClient.SelectContext(ctx, &folders, sql, args...)
	if err != nil {
		return nil, err
	}

	_folders := []*entity.Folder{}
	for i := range folders {
		_folders = append(_folders, &folders[i])
	}
	return _folders, nil
}

func (repo *classRepository) PostFolders(ctx context.Context, classID int64, name string) error {
	now := time.Now()
	folder := entity.Folder{
		ClassID:   classID,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	sql := "INSERT INTO folders (class_id, name, created_at, updated_at) VALUES (:class_id, :name, :created_at, :updated_at)"
	_, err := repo.DBClient.NamedExecContext(ctx, sql, folder)
	if err != nil {
		return err
	}
	return nil
}
