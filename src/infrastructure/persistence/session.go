package persistence

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gomodule/redigo/redis"
	"github.com/hashmup/QuestionBankAPI/src/config"
	"github.com/hashmup/QuestionBankAPI/src/domain/entity"
	"github.com/hashmup/QuestionBankAPI/src/domain/repository"
	"github.com/jmoiron/sqlx"
)

type sessionRepository struct {
	DBClient    *sqlx.DB
	RedisClient redis.Conn
}

func NewSessionRepository(dbc *sqlx.DB, rc redis.Conn) repository.SessionRepository {
	return &sessionRepository{
		DBClient:    dbc,
		RedisClient: rc,
	}
}

func (repo *sessionRepository) CreateSession(ctx context.Context, userID int64) (*entity.Session, error) {
	token, err := config.GenerateToken(20)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	session := entity.Session{
		UserID:    userID,
		Token:     token,
		ExpiredAt: now.AddDate(0, 1, 0),
		CreatedAt: now,
		UpdatedAt: now,
	}
	sql := "INSERT INTO sessions (user_id, token, expired_at, created_at, updated_at) VALUES (:user_id, :token, :expired_at, :created_at, :updated_at)"

	// Insert a row to AuthToken table
	res, err := repo.DBClient.NamedExecContext(ctx, sql, session)
	if err != nil {
		panic(err)
	}

	session.ID, _ = res.LastInsertId()

	repo.RedisClient.Do("SET", userID, token)
	repo.RedisClient.Do("EXPIRE", userID, config.RedisTTL)
	return &session, nil
}

func (repo *sessionRepository) DeleteSession(ctx context.Context, userID int64, token string) (bool, error) {
	sql, args, _ := sq.Update("sessions").Set("deleted_at", time.Now()).Where(sq.Eq{"user_id": userID, "token": token}).ToSql()
	_, err := repo.DBClient.ExecContext(ctx, sql, args...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *sessionRepository) IsValidSession(ctx context.Context, userID int64, token string) (bool, error) {
	session := entity.Session{}
	sql, args, _ := sq.Select("*").From("sessions").Where(sq.Eq{"token": token, "user_id": userID, "deleted_at": nil}).ToSql()
	fmt.Printf("%s\n", sql)
	err := repo.DBClient.GetContext(ctx, &session, sql, args...)
	if err != nil {
		return false, nil
	}
	if session.ExpiredAt.Before(time.Now()) {
		// This means expiration date < time.now, thus this token is already expired
		// Notify client to re-login
		return false, nil
	}
	// o/w update the expiration date for both DB and redis
	repo.updateExpirationDate(ctx, session)
	return true, nil
}

func (repo *sessionRepository) updateExpirationDate(ctx context.Context, session entity.Session) {
	session.ExpiredAt = time.Now().AddDate(0, 1, 0)
	sql := "UPDATE sessions SET expired_at = :expired_at WHERE id = :id"
	_, err := repo.DBClient.NamedExecContext(ctx, sql, session)
	if err != nil {
		panic(err)
	}
	repo.RedisClient.Do("EXPIRE", session.UserID, config.RedisTTL)
}
