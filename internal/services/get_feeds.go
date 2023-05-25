package services

import (
	"context"
	"database/sql"
	"fmt"
)

func GetFeeds(ctx context.Context, DBConn *sql.DB) ([]Feed, error) {
	feeds := []Feed{}
	query := "SELECT * FROM feeds"
	rows, err := DBConn.QueryContext(ctx, query)
	if err != nil {
		return feeds, fmt.Errorf("no feeds found: %v", err)
	}

	for rows.Next() {
		feed := Feed{}
		err := rows.Scan(&feed.ID, &feed.CreatedAt, &feed.UpdatedAt, &feed.Name, &feed.URL, &feed.UserID)
		if err != nil {
			return feeds, fmt.Errorf("error getting feeds: %v", err)
		}
		feeds = append(feeds, feed)
	}
	return feeds, nil
}
