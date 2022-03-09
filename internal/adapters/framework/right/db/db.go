package db

import (
	"log"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Adapter struct {
	db *sqlx.DB
}

func NewAdapter(dName, dsName string) (*Adapter, error) {
	db, err := sqlx.Open(dName, dsName)
	if err != nil {
		log.Printf("Could open db: %v", err)
		os.Exit(1)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error while checking db PING: %v", err)
		return nil, err
	}

	return &Adapter{db}, nil
}

func (ad Adapter) CloseDBConntection() error {
	return ad.db.Close()
}

func (ad Adapter) AddToHistory(answer int32, operation string) error {
	qs, args, err := squirrel.Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		log.Fatalf("Error while inserting to db: %v", err)
	}

	_, err = ad.db.Exec(qs, args...)
	if err != nil {
		log.Fatalf("Error while executing query string: %v", err)
	}

	return nil
}
