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

func BlogView(auth views.IAuth) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		db := db.Open()
		defer db.Close()

		stag := views.PathValue(ctx, "tag")

		if (stag != &views.EmptyString) && !db.Query(&SQL_CHECK_TAGS, auth.State(), stag).Get() {
			return http_errors.NotFoundView(auth)(ctx)
		}

		data := Blog{
			stag: stag,

			tags:     QueryTags(db, auth, stag),
			articles: QueryArticles(db, auth, stag),
		}

		views.WritePage(ctx, &data, auth)

		return nil
	}
}
