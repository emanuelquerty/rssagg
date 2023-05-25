package services

import (
	"context"
	"database/sql"
)

func GetNextFeedToFetch(ctx context.Context, DBConn *sql.DB) (Feed, error) {
	feed := Feed{}
	query :=
		`SELECT * FROM feeds
	ORDER BY last_fetched_at ASC NULLS FIRST
	LIMIT 1`

	rows := DBConn.QueryRowContext(ctx, query)
	err := rows.Scan(&feed.ID, &feed.CreatedAt, &feed.UpdatedAt,
		&feed.Name, &feed.URL, &feed.UserID, &feed.LastFetchedAt)

	if err != nil {
		return feed, err
	}

	return feed, nil
}
