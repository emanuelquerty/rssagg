package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"  bson:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"  bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"  bson:"updated_at"`
	Name      string    `json:"name,omitempty"  bson:"name"`
	ApiKey    string    `json:"api_key,omitempty"  bson:"api_key"`
}

func CreateUser(ctx context.Context, DBConn *sql.DB, user User) (User, error) {
	query := `
	INSERT INTO 
	users(id, created_at, updated_at, name, api_key) 
	VALUES($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex')) 
	RETURNING *`

	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Could not prepare statement %v", err)
		return User{}, err
	}
	row := stmt.QueryRowContext(ctx, user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	newlyCreatedUser := User{}
	err = row.Scan(&newlyCreatedUser.ID, &newlyCreatedUser.CreatedAt,
		&newlyCreatedUser.UpdatedAt, &newlyCreatedUser.Name, &newlyCreatedUser.ApiKey)
	if err != nil {
		log.Printf("Could not created a new user %v", err)
		return newlyCreatedUser, err
	}

	return newlyCreatedUser, nil
}
