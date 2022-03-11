package db

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(dName, dSource string) (*DB, error) {
	nd, err := sqlx.Open(dName, dSource)
	if err != nil {
		return nil, err
	}

	if err = nd.Ping(); err != nil {
		return nil, err
	}

	return &DB{db: nd}, nil
}

func (eb *DB) CloseConnection() error {
	return eb.db.Close()
}

func (eb *DB) AddToHistory(value int32, operation string) error {

	qs, args, err := squirrel.Insert("arith_history").
		Columns("date", "value", "operation").
		Values(time.Now(), value, operation).ToSql()
	if err != nil {
		return err
	}

	if _, err = eb.db.Exec(qs, args...); err != nil {
		return err
	}

	return nil
}
