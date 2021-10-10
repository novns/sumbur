package blog

import (
	_ "embed"
	"sumbur/db"
	"sumbur/views"
	"sumbur/views/http_errors"

	"github.com/savsgio/atreugo/v11"
)

//go:embed sql/tag.sql
var SQL_CHECK_TAGS string

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
