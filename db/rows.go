package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Rows struct {
	ctx  *context.Context
	conn *pgxpool.Pool
	rows pgx.Rows

	query *string
	args  *[]interface{}
	dest  *[]interface{}
}

func (rows *Rows) Fields(dest ...interface{}) *Rows {
	rows.dest = &dest
	return rows
}

func (rows *Rows) Next() bool {
	var err error

	if rows.rows == nil {
		rows.rows, err = rows.conn.Query(*rows.ctx, *rows.query, *rows.args...)
		if err != nil {
			panic(err)
		}
	}

	if !rows.rows.Next() {
		rows.Close()

		return false
	}

	err = rows.rows.Scan(*rows.dest...)
	if err != nil {
		panic(err)
	}

	return true
}

func (rows *Rows) Close() {
	if rows.rows != nil {
		rows.rows.Close()
	}

	rows.rows = nil
	rows.args = nil
	rows.dest = nil
}
