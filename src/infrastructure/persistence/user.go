package persistence

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	DBClient *sqlx.DB
}

func NewUserRepository(dbc *sqlx.DB) repository.UserRepository {
	return &userRepository{
		DBClient: dbc,
	}
}

func (repo *userRepository) PostUserRegister(ctx context.Context, name, email, password, accountType string, schoolID int64) (*entity.UserResponse, error) {
	now := time.Now()
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := entity.User{
		Name:              name,
		Email:             email,
		EncryptedPassword: string(encryptedPassword),
		SchoolID:          schoolID,
		Type:              accountType,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
	sql := "INSERT INTO users (name, email, encrypted_password, school_id, type, created_at, updated_at) VALUES (:name, :email, :encrypted_password, :school_id, :type, :created_at, :updated_at)"

	// Insert a row to AuthToken table
	res, err := repo.DBClient.NamedExecContext(ctx, sql, user)
	if err != nil {
		return nil, err
	}

	user.ID, _ = res.LastInsertId()
	school := entity.School{}
	sql, args, _ := sq.Select("*").From("schools").Where(sq.Eq{"id": user.SchoolID}).ToSql()
	err = repo.DBClient.GetContext(ctx, &school, sql, args...)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		SchoolID:   user.SchoolID,
		SchoolName: school.Name,
		Type:       user.Type,
	}, nil
}

func (repo *userRepository) UserLogin(ctx context.Context, email, password string) (*entity.User, error) {
	user := entity.User{}
	sql, args, _ := sq.Select("*").From("users").Where(sq.Eq{"email": email}).ToSql()
	err := repo.DBClient.GetContext(ctx, &user, sql, args...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) GetUser(ctx context.Context, userID int64) (*entity.UserResponse, error) {
	user := entity.User{}
	sql, args, _ := sq.Select("*").From("users").Where(sq.Eq{"id": userID}).ToSql()
	err := repo.DBClient.GetContext(ctx, &user, sql, args...)
	if err != nil {
		return nil, err
	}
	school := entity.School{}
	sql, args, _ = sq.Select("*").From("schools").Where(sq.Eq{"id": user.SchoolID}).ToSql()
	err = repo.DBClient.GetContext(ctx, &school, sql, args...)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		SchoolID:   user.SchoolID,
		SchoolName: school.Name,
		Type:       user.Type,
	}, nil
}
