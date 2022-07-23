package db

import (
	"context"
	"database-interface/db/sqlc"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB      *sql.DB
	CTX     context.Context
	Queries *sqlc.Queries
)

func Run() {
	var err error
	CTX = context.Background()

	DB, err = sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONSTRING"))
	if err != nil {
		fmt.Println(err)
	}

	Queries = sqlc.New(DB)

}
