package services

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func GetPosts(ctx context.Context, DBConn *sql.DB, userID uuid.UUID, postsCount int) ([]Post, error) {
	posts := []Post{}
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
		post := Post{}

		err := rows.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title,
			&post.Description, &post.PublishedAt, &post.Url, &post.FeedId)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
