package entity

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID                int64          `db:"id" json:"id"`
	Name              string         `db:"name" json:"name"`
	Email             string         `db:"email" json:"email"`
	EncryptedPassword string         `db:"encrypted_password" json:"-"`
	SchoolID          int64          `db:"school_id" json:"school_id"`
	Type              string         `db:"type" json:"type"`
	CreatedAt         time.Time      `db:"created_at" json:"-"`
	UpdatedAt         time.Time      `db:"updated_at" json:"-"`
	DeletedAt         mysql.NullTime `db:"deleted_at" json:"-"`
}

type UserResponse struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	SchoolID   int64  `json:"school_id"`
	SchoolName string `json:"school_name"`
	Type       string `json:"type"`
}
