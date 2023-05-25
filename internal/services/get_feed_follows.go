package services

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func GetFeedFollows(ctx context.Context, DBConn *sql.DB, userID uuid.UUID) ([]FeedFollows, error) {
	feedFollows := []FeedFollows{}
	query := `SELECT * FROM feed_follows WHERE user_id=$1`

	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		return feedFollows, err
	}

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return feedFollows, err
	}

	for rows.Next() {
		feedFollow := FeedFollows{}

		err = rows.Scan(&feedFollow.ID, &feedFollow.CreatedAt,
			&feedFollow.UpdatedAt, &feedFollow.UserID, &feedFollow.FeedID)
		if err != nil {
			return feedFollows, err
		}

		feedFollows = append(feedFollows, feedFollow)
	}
	return feedFollows, nil
}
