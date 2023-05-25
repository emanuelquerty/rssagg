package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func DeleteFeedFollows(ctx context.Context, DBConn *sql.DB, ID uuid.UUID, userID uuid.UUID) error {
	query := "DELETE FROM feed_follows WHERE id = $1 AND user_id = $2"

	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, ID, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("cannot unfollow feed you are not currently following: %v", rows)
	}
	return nil
}
