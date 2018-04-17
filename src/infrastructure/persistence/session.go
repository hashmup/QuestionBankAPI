package persistence

import (
	"context"
	"time"

	"github.com/Graffity-X/user-api/src/config"
	"github.com/gomodule/redigo/redis"
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
