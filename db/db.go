package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Connection

var DSN *string

type DB struct {
	ctx  *context.Context
	conn *pgxpool.Pool
}

func Open(ctx context.Context) *DB {
	conn, err := pgxpool.Connect(ctx, *DSN)
	if err != nil {
		panic(err)
	}

	return &DB{ctx: &ctx, conn: conn}
}

func (db *DB) Close() {
	db.conn.Close()
}

// Transactions

func (db *DB) Begin() {
	db.Exec("begin")
}

func (db *DB) Commit() {
	db.Exec("commit")
}

// Queries

func (db *DB) Exec(query string, args ...interface{}) int64 {
	ct, err := db.conn.Exec(*db.ctx, query, args...)
	if err != nil {
		panic(err)
	}

	return ct.RowsAffected()
}

func (db *DB) Query(query *string, args ...interface{}) *Rows {
	return &Rows{
		ctx:   db.ctx,
		conn:  db.conn,
		query: query,
		args:  &args,
	}
}
