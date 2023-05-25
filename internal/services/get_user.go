package services

import (
	"context"
	"database/sql"
	"log"
)

func GetUser(ctx context.Context, DBConn *sql.DB, apiKey string) (User, error) {
	query := "SELECT * FROM users WHERE api_key=$1"
	stmt, err := DBConn.PrepareContext(ctx, query)
	if err != nil {
		log.Println("Auth Error: error preparing statement", err)
		return User{}, err
	}

	row := stmt.QueryRowContext(ctx, apiKey)

	user := User{}
	err = row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.ApiKey)
	if err != nil {
		log.Println("ErrNoRows: No user found", err)
		return User{}, err
	}
	return user, nil
}
