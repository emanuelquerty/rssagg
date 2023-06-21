package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID `json:"id,omitempty"  bson:"id"`
	CreatedAt     time.Time `json:"created_at,omitempty"  bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"  bson:"updated_at"`
	Name          string    `json:"name,omitempty"  bson:"name"`
	ApiKey        string    `json:"api_key,omitempty"  bson:"api_key"`
	URL           string    `json:"url,omitempty"  bson:"url"`
	UserID        uuid.UUID `json:"user_id,omitempty"  bson:"user_id"`
	LastFetchedAt time.Time `json:"last_fetched_at,omitempty"  bson:"last_fetched_at"`
}

func CreateFeed(ctx context.Context, DBConn *sql.DB, feed Feed) (Feed, error) {
	query := `
	INSERT INTO 
	feeds(id, created_at, updated_at, name, url, user_id, last_fetched_at) 
	VALUES($1, $2, $3, $4, $5, $6, $7) 
	RETURNING *`

	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Could not prepare statement %v", err)
		return Feed{}, err
	}
	row := stmt.QueryRowContext(ctx, feed.ID, feed.CreatedAt,
		feed.UpdatedAt, feed.Name, feed.URL, feed.UserID, feed.LastFetchedAt)

	newlyCreatedFeed := Feed{}

	err = row.Scan(&newlyCreatedFeed.ID, &newlyCreatedFeed.CreatedAt,
		&newlyCreatedFeed.UpdatedAt, &newlyCreatedFeed.Name,
		&newlyCreatedFeed.URL, &newlyCreatedFeed.UserID,
		&newlyCreatedFeed.LastFetchedAt)

	if err != nil {
		log.Printf("Could not created a new feed %v", err)
		return newlyCreatedFeed, err
	}

	return newlyCreatedFeed, nil
}
