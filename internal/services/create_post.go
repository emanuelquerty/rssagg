package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID      `json:"id,omitempty"  bson:"id"`
	CreatedAt   time.Time      `json:"created_at,omitempty"  bson:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"  bson:"updated_at"`
	Title       string         `json:"title,omitempty"  bson:"title"`
	Description sql.NullString `json:"description,omitempty"  bson:"description"`
	PublishedAt time.Time      `json:"published_at,omitempty"  bson:"published_at"`
	Url         string         `json:"url,omitempty"  bson:"url"`
	FeedId      uuid.UUID      `json:"feed_id,omitempty"  bson:"feed_id"`
}

func (sctx *ServiceContext) CreatePost(ctx context.Context, post Post) error {
	query :=
		`INSERT INTO posts
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := sctx.DBConn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, post.ID, post.CreatedAt, post.UpdatedAt,
		post.Title, post.Description, post.PublishedAt, post.Url, post.FeedId)
	if err != nil {
		return err
	}

	rowsAffectedCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffectedCount != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %v rows instead", rowsAffectedCount)
	}
	return nil
}
