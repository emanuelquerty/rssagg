package handlers

import (
	"database/sql"
	"log"
)

type HandlerContext struct {
	DBConn *sql.DB
	Logger *log.Logger
}
