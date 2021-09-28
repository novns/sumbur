package http_errors

import (
	"runtime/debug"
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type Panic struct {
	*views.BasePage

	err   interface{}
	stack []byte
}

func PanicView(auth views.IAuth) atreugo.PanicView {
	return func(ctx *atreugo.RequestCtx, err interface{}) {
		ctx.Logger().Printf("%s", err)

		ctx.SetStatusCode(500)
		views.WritePage(ctx, &Panic{err: err, stack: debug.Stack()}, auth)
	}
}
