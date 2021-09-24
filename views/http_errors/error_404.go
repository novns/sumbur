package http_errors

import (
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type NotFound struct {
	*views.BasePage

	path []byte
}

func GetNotFound(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(404)
	views.WritePage(ctx, &NotFound{path: ctx.Path()})

	return nil
}
