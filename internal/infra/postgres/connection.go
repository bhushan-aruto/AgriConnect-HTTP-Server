package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type connection struct {
	Pool *pgxpool.Pool
}

func NewDatabase(databaseUrl string) *connection {
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalln("error occurred while connecting to database, Error: ", err.Error())
	}

	log.Println("connected to the database with pooling")

	return &connection{
		Pool: pool,
	}

}

func (db *connection) CheckDatabase() {
	if err := db.Pool.Ping(context.Background()); err != nil {
		log.Fatalln("error occurred while performing database healthcheck, Error: ", err.Error())
	}

	log.Println("database was working correctly")
}

func (db *connection) CloseConnection() {
	db.Pool.Close()
}
