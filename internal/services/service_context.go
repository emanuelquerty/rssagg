package services

import (
	"database/sql"
	"log"
)

type ServiceContext struct {
	Logger log.Logger
	DBConn *sql.DB
}

func NewServiceContext(logger log.Logger, dbConn *sql.DB) ServiceContext {
	ctx := ServiceContext{
		Logger: logger,
		DBConn: dbConn,
	}
	return ctx
}
