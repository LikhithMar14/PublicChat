package store

import (
	"context"
	"database/sql"

	"github.com/LikhithMar14/social/internal/models"
)

type User struct{
	db *sql.DB
}

func (s *User) Create(ctx context.Context, user *models.User) error{
	query := `INSERT INTO users (email,username,password)
			 VALUES ($1,$2,$3)
			 RETURNING id,created_at,updated_at`

	row := s.db.QueryRowContext(ctx,query,user.Email,user.Username,user.Password)
	
	err := row.Scan(&user.ID,&user.CreatedAt,&user.UpdatedAt)
	if err != nil{
		return err
	}
	return nil
}