package views

import (
	"net/url"
	"strconv"

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

func IntValue(ctx *atreugo.RequestCtx, key string) int {
	value := ctx.UserValue(key)
	if value == nil {
		return 0
	}

	result, err := strconv.Atoi(value.(string))
	if err != nil {
		return 0
	}

	return result
}
