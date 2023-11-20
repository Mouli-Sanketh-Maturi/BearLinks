package datastore

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var db *pgxpool.Pool

func GetDb() *pgxpool.Pool {
	if db == nil {
		var err error
		db, err = pgxpool.New(ctx, "postgresql://moulisanketh:password@localhost/bearLinks")

		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}
