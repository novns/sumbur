package views

import (
	"net/url"

	"github.com/savsgio/atreugo/v11"
)

var EmptyString = ""

func PathValue(ctx *atreugo.RequestCtx, key string) *string {
	value := ctx.UserValue(key)
	if value == nil {
		return &EmptyString
	}

	result, _ := url.PathUnescape(value.(string))
	return &result
}
