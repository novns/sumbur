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

	RowsAffected int64
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

func (db *DB) Begin() *DB {
	db.conn.Exec(*db.ctx, "begin")
	return db
}

func (db *DB) Commit() *DB {
	db.conn.Exec(*db.ctx, "commit")
	return db
}

// Queries

func (db *DB) Exec(query string, args ...interface{}) *DB {
	ct, err := db.conn.Exec(*db.ctx, query, args...)
	if err != nil {
		panic(err)
	}

	db.RowsAffected = ct.RowsAffected()

	return db
}

func (db *DB) Query(query *string, args ...interface{}) *Rows {
	db.RowsAffected = 0

	return &Rows{
		ctx:   db.ctx,
		conn:  db.conn,
		query: query,
		args:  &args,
	}
}
