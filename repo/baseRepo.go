package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type IRepository interface {
	Connect() (*sqlx.DB, error)
	Create(ctx context.Context, query string, dest interface{}, params ...interface{}) (interface{}, error)
	FindById(ctx context.Context, query string, dest interface{}, params ...interface{}) (interface{}, error)
	Update(e any)
	Delete(e any)
}

type AbstractRepo struct {
	DriverName string
	IRepository
}
type PGError struct {
	Code    string
	Message string
}

func (r *AbstractRepo) Connect() (*sqlx.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOSTNAME"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Open(r.DriverName, connectionString)
	if err != nil {
		//log error
		return nil, err
	}

	if err := db.Ping(); err != nil {
		//log the error
		return nil, err
	}

	return db, nil
}

func (r *AbstractRepo) Create(ctx context.Context, query string, dest interface{}, params ...interface{}) error {
	db, err := r.Connect()
	if err != nil {
		return err
	}
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	err = tx.QueryRowxContext(ctx, query, params...).Scan(dest)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			fmt.Println(pgErr.Code)
			errors.New("")
		}
		tx.Rollback()
		return err
	}
	tx.Commit()
	db.Close()
	return nil
}

func (r *AbstractRepo) FindById(query string, dest interface{}, id int) error {
	err := r.RunQuery(query, dest, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AbstractRepo) RunQuery(query string, dest interface{}, params ...interface{}) error {
	db, err := r.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.Get(dest, params...)
	if err != nil {
		return err
	}

	return nil
}
