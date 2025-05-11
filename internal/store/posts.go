package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LikhithMar14/social/internal/models"
	"github.com/lib/pq"
)

type Post struct {
	db *sql.DB
}

var (
	ErrNotFound = errors.New("records not found	")
)

func (s *Post) Create(ctx context.Context, post *models.Post) error {

	query := `INSERT INTO posts (content,title,user_id,tags)
			 VALUES ($1,$2,$3,$4)
			 RETURNING id,created_at,updated_at`

	row := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags))

	err := row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Post)GetByID(ctx context.Context, id int64)(*models.Post,error){
	query := `SELECT id,user_id,content,title,tags,created_at,updated_at FROM posts WHERE id = $1`

	row := s.db.QueryRowContext(ctx, query, id)
	var post models.Post
	err := row.Scan(&post.ID, &post.UserID, &post.Content, &post.Title, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt)
	if err != nil{
		switch{
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &post, nil
	
}
