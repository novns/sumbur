package blog

import (
	_ "embed"
	"sumbur/db"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

//go:embed sql/tags.sql
var SQL_TAGS string

type Tags struct {
	query *db.Rows

	stag *string

	tag   []byte
	count int
}

func QueryTags(db *db.DB, stag *string) *Tags {
	tags := Tags{
		query: db.Query(SQL_TAGS, views.AuthState),
		stag:  stag,
	}

	tags.query.Fields(
		&tags.tag,
		&tags.count,
	)

	return &tags
}

func (tags *Tags) Next() bool {
	return tags.query.Next()
}

func TagsGet(ctx *atreugo.RequestCtx) error {
	db := db.Open(ctx)
	defer db.Close()

	QueryTags(db, views.PathValue(ctx, "tag")).WriteTags(ctx)

	return nil
}
