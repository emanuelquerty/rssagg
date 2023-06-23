package handlers

import (
	"database/sql"
	"log"
)

type HandlerContext struct {
	DBConn *sql.DB
	Logger *log.Logger
}

func NewHandlerContext(dbConn *sql.DB, logger *log.Logger) HandlerContext {
	hctx := HandlerContext{
		DBConn: dbConn,
		Logger: logger,
	}
	return hctx
}
