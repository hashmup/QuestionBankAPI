package persistence

import (
	"context"
	"errors"
	"fmt"
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

func (repo *classRepository) GetClasses(ctx context.Context, schoolID int64, name, classCode, term string) ([]*entity.Class, error) {
	query := sq.Select("*").From("classes").Where(sq.Eq{"school_id": schoolID})
	if name != "" {
		query = query.Where(sq.Expr("name LIKE ?", "%"+name+"%"))
	}
	if classCode != "" {
		query = query.Where(sq.Eq{"class_code": classCode})
	}
	if term != "" {
		query = query.Where(sq.Eq{"term": term})
	}
	sql, args, _ := query.ToSql()
	fmt.Printf("%s %#v\n", sql, args)
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

func (repo *classRepository) GetClassesBySchoolID(ctx context.Context, schoolID int64) ([]*entity.Class, error) {
	sql, args, _ := sq.Select("*").From("classes").Where(sq.Eq{"school_id": schoolID}).ToSql()
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

func (repo *classRepository) GetClassesByUserID(ctx context.Context, userID int64) ([]*entity.Class, error) {
	sql, args, _ := sq.Select("classes.*").From("user_class_assoc").Join("classes ON user_class_assoc.class_id = classes.id").Where(sq.Eq{"user_class_assoc.user_id": userID}).ToSql()
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

func (repo *classRepository) PostFolders(ctx context.Context, classID int64, name, description string) error {
	now := time.Now()
	folder := entity.Folder{
		ClassID:     classID,
		Name:        name,
		Description: &description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	sql := "INSERT INTO folders (class_id, name, description, created_at, updated_at) VALUES (:class_id, :name, :description, :created_at, :updated_at)"
	_, err := repo.DBClient.NamedExecContext(ctx, sql, folder)
	if err != nil {
		return err
	}
	return nil
}

func (repo *classRepository) JoinClass(ctx context.Context, userID, classID int64) error {
	userClassAssoc := entity.UserClassAssoc{}
	sql, args, _ := sq.Select("*").From("user_class_assoc").Where(sq.Eq{"user_id": userID, "class_id": classID}).ToSql()
	err := repo.DBClient.GetContext(ctx, &userClassAssoc, sql, args...)
	if err != nil {
		return errors.New("User already joined the class")
	}
	now := time.Now()
	userClassAssoc = entity.UserClassAssoc{
		UserID:    userID,
		ClassID:   classID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	sql = "INSERT INTO user_class_assoc (user_id, class_id, created_at, updated_at) VALUES (:user_id, :class_id, :created_at, :updated_at)"
	_, err = repo.DBClient.NamedExecContext(ctx, sql, userClassAssoc)
	if err != nil {
		return err
	}
	return nil
}

func (repo *classRepository) CreateClass(ctx context.Context, userID int64, name, code, term string) error {
	user := entity.User{}
	sql, args, _ := sq.Select("*").From("users").Where(sq.Eq{"id": userID}).ToSql()
	err := repo.DBClient.GetContext(ctx, &user, sql, args...)
	if err != nil {
		return err
	}
	now := time.Now()
	class := entity.Class{
		Name:         name,
		SchoolID:     user.SchoolID,
		InstructorID: user.ID,
		Term:         term,
		ClassCode:    code,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	sql = "INSERT INTO classes (name, school_id, instructor_id, term, class_code, created_at, updated_at) VALUES (:name, :school_id, :instructor_id, :term, :class_code, :created_at, :updated_at)"
	res, err := repo.DBClient.NamedExecContext(ctx, sql, class)
	if err != nil {
		return err
	}

	now = time.Now()
	classID, _ := res.LastInsertId()
	userClassAssoc := entity.UserClassAssoc{
		UserID:    userID,
		ClassID:   classID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	sql = "INSERT INTO user_class_assoc (user_id, class_id, created_at, updated_at) VALUES (:user_id, :class_id, :created_at, :updated_at)"
	_, err = repo.DBClient.NamedExecContext(ctx, sql, userClassAssoc)
	if err != nil {
		return err
	}
	return nil
}
