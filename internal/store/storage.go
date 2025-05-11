package store

import (
	"context"
	"database/sql"

	"github.com/LikhithMar14/social/internal/models"
)

type Storage struct{
	Posts interface{
		Create(context.Context, *models.Post) error
		GetByID(context.Context, int64)(*models.Post,error)
	}
	Users interface{
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *sql.DB) Storage{
	return Storage{
		Posts: &Post{db: db},
		Users: &User{db: db},
	}
}