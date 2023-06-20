package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PostWithNullDescriptionAllowed struct {
	ID          uuid.UUID `json:"id,omitempty"  bson:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty"  bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"  bson:"updated_at"`
	Title       string    `json:"title,omitempty"  bson:"title"`
	Description *string   `json:"description,omitempty"  bson:"description"`
	PublishedAt time.Time `json:"published_at,omitempty"  bson:"published_at"`
	Url         string    `json:"url,omitempty"  bson:"url"`
	FeedId      uuid.UUID `json:"feed_id,omitempty"  bson:"feed_id"`
}

func GetPosts(ctx context.Context, DBConn *sql.DB, userID uuid.UUID, postsCount int) ([]PostWithNullDescriptionAllowed, error) {
	posts := []PostWithNullDescriptionAllowed{}
	query := `
	SELECT posts.* FROM posts
	JOIN feed_follows 
	ON posts.feed_id = feed_follows.feed_id
	WHERE feed_follows.user_id = $1
	ORDER BY posts.published_at DESC
	LIMIT $2`

	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		return posts, err
	}

	rows, err := stmt.QueryContext(ctx, userID, postsCount)
	if err != nil {
		return posts, err
	}

	for rows.Next() {
		post := PostWithNullDescriptionAllowed{}

		err := rows.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title,
			&post.Description, &post.PublishedAt, &post.Url, &post.FeedId)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
