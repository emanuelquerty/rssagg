package services

import (
	"context"

	"github.com/google/uuid"
)

func (sctx *ServiceContext) GetNextFeedsToFetch(ctx context.Context, feedCount int) ([]Feed, error) {
	feeds := []Feed{}
	query :=
		`SELECT id FROM feeds
		ORDER BY last_fetched_at ASC NULLS FIRST 
		LIMIT $1`

	stmt, err := sctx.DBConn.PrepareContext(ctx, query)
	if err != nil {
		return feeds, err
	}

	rows, err := stmt.QueryContext(ctx, feedCount)
	if err != nil {
		return feeds, err
	}

	for rows.Next() {
		feed := Feed{}
		err := rows.Scan(&feed.ID)
		if err != nil {
			return feeds, err
		}
		feed, err = sctx.MarkFeedAsFetched(ctx, feed.ID)
		if err != nil {
			return feeds, err
		}
		feeds = append(feeds, feed)
	}
	return feeds, nil
}

func (sctx *ServiceContext) MarkFeedAsFetched(ctx context.Context, feedID uuid.UUID) (Feed, error) {
	feed := Feed{}
	query :=
		`UPDATE feeds 
		SET 
		last_fetched_at = NOW(), 
		updated_at = NOW() 
		WHERE id = $1 
		RETURNING *`

	stmt, err := sctx.DBConn.PrepareContext(ctx, query)
	if err != nil {
		return feed, err
	}

	row := stmt.QueryRowContext(ctx, feedID)
	err = row.Scan(&feed.ID, &feed.CreatedAt, &feed.UpdatedAt,
		&feed.Name, &feed.URL, &feed.UserID, &feed.LastFetchedAt)

	if err != nil {
		return feed, err
	}
	return feed, nil
}
