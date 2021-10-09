package blog

import (
	"sumbur/db"
	"sumbur/views"
	"sumbur/views/http_errors"

	"github.com/MakeNowJust/heredoc"
	"github.com/savsgio/atreugo/v11"
)

var SQL_CHECK_TAGS = heredoc.Doc(`
SELECT "tag"."tag"
FROM "tag"
INNER JOIN "article"  USING ("article_id")
WHERE
	($1 OR "article"."public") AND ("tag"."tag" = $2)
LIMIT 1
`)

type Blog struct {
	*views.BasePage

	stag *string

	tags     *Tags
	articles *Articles
}

func BlogGet(ctx *atreugo.RequestCtx) error {
	db := db.Open(ctx)
	defer db.Close()

	stag := views.PathValue(ctx, "tag")

	if (stag != &views.EmptyString) && !db.Query(SQL_CHECK_TAGS, views.AuthState, stag).Get() {
		return http_errors.NotFoundView(ctx)
	}

	data := Blog{
		stag: stag,

		tags:     QueryTags(db, stag),
		articles: QueryArticles(db, stag),
	}

	views.WritePage(ctx, &data)

	return nil
}
