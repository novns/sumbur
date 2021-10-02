package blog

import (
	"sumbur/db"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type Blog struct {
	*views.BasePage
}

func BlogView(auth views.IAuth) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		db := db.Open()
		defer db.Close()

		data := Blog{}

		views.WritePage(ctx, &data, auth)

		return nil
	}
}
