package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type PgConnectionManager struct {
	Db *sqlx.DB
}

func NewPgConnectionManager() (*PgConnectionManager, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOSTNAME"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		//log error
		return nil, err
	}

	if err := db.Ping(); err != nil {
		//log the error
		return nil, err
	}

	return &PgConnectionManager{
		Db: db,
	}, nil
}
