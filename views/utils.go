package views

import (
	"net/url"

	"github.com/savsgio/atreugo/v11"
)

func PathStringValue(ctx *atreugo.RequestCtx, key string) string {
	value := ctx.UserValue(key)
	if value == nil {
		return ""
	}

	result, _ := url.PathUnescape(value.(string))
	return result
}
