package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FeedFollows struct {
	ID        uuid.UUID `json:"id,omitempty"  bson:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"  bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"  bson:"updated_at"`
	UserID    uuid.UUID `json:"user_id,omitempty"  bson:"user_id"`
	FeedID    uuid.UUID `json:"feed_id,omitempty"  bson:"feed_id"`
}

func CreateFeedFollows(ctx context.Context, DBConn *sql.DB, feedFollows FeedFollows) error {
	query :=
		`INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
	VALUES($1, $2, $3, $4, $5)
	RETURNING *`

	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, feedFollows.ID, feedFollows.CreatedAt,
		feedFollows.UpdatedAt, feedFollows.UserID, feedFollows.FeedID)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rowsCount)
	}
	return nil
}
