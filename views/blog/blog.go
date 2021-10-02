package blog

import (
	"sumbur/db"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type Blog struct {
	*views.BasePage

	tags     *Tags
	articles *Articles
}

func BlogView(auth views.IAuth) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		db := db.Open()
		defer db.Close()

		tag := views.PathValue(ctx, "tag")

		data := Blog{
			tags:     QueryTags(db, auth, tag),
			articles: QueryArticles(db, auth, tag),
		}

		views.WritePage(ctx, &data, auth)

		return nil
	}
}
