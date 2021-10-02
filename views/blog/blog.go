package blog

import (
	"sumbur/db"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type Blog struct {
	*views.BasePage

	tags *Tags
}

func BlogView(auth views.IAuth) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		db := db.Open()
		defer db.Close()

		tag := views.PathStringValue(ctx, "tag")

		data := Blog{
			tags: QueryTags(db, auth, &tag),
		}

		views.WritePage(ctx, &data, auth)

		return nil
	}
}
