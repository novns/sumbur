package blog

import (
	"sumbur/db"
	"sumbur/views"

	"github.com/MakeNowJust/heredoc"
	"github.com/savsgio/atreugo/v11"
)

var SQL_TAGS = heredoc.Doc(`
SELECT
	"tag",
	count("article_id")  AS "cnt"
FROM "tag"
INNER JOIN "article"  USING ("article_id")
WHERE
	$1 OR "article"."public"
GROUP BY
	"tag"
ORDER BY
	"cnt" DESC,
	"tag"
`)

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
