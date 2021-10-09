package blog

import (
	"sumbur/db"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func EditGet(ctx *atreugo.RequestCtx) error {
	db := db.Open(ctx)
	defer db.Close()

	article := &Article{stag: &views.EmptyString}
	article_id := views.IntValue(ctx, "article_id")

	if article.QueryArticle(db, article_id) {
		article.WriteEdit(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}

	return nil
}
