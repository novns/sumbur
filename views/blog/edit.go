package blog

import (
	"bytes"
	"sumbur/db"
	"sumbur/views"

	"github.com/MakeNowJust/heredoc"
	"github.com/gomarkdown/markdown"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

//

func EditGet(ctx *atreugo.RequestCtx) error {
	db := db.Open(ctx)
	defer db.Close()

	article := &Article{stag: &views.EmptyString}
	article_id := views.IntValue(ctx, "article_id")

	if article_id == 0 {
		article.WriteEdit(ctx)
		return nil
	}

	if article.QueryArticle(db, article_id) {
		article.WriteEdit(ctx)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}

	return nil
}

//

var SQL_INSERT = heredoc.Doc(`
INSERT INTO "article" (
	"title",
	"body",
	"body_ht",
	"public"
)
VALUES (
	$1,
	$2,
	$3,
	$4
)
ON CONFLICT  DO NOTHING
RETURNING "article_id"
`)

var SQL_UPDATE = heredoc.Doc(`
UPDATE "article"
SET
	"title" = $2,
	"body" = $3,
	"body_ht" = $4,
	"public" = $5
WHERE "article_id" = $1
`)

func EditPost(ctx *atreugo.RequestCtx) error {
	db := db.Open(ctx)
	defer db.Close()

	// Article

	article_id := views.IntValue(ctx, "article_id")

	args := ctx.PostArgs()
	body := args.Peek("body")
	body_ht := []byte{}

	if len(body) > 0 {
		body = bytes.ReplaceAll(body, []byte("\r"), []byte{})
		body_ht = markdown.ToHTML(body, nil, nil)
	}

	db.Begin()

	if article_id == 0 {
		db.Query(
			SQL_INSERT,
			args.Peek("title"),
			body,
			body_ht,
			args.Has("public"),
		).Get(&article_id)

		views.SetIntValue(ctx, "article_id", article_id)
	} else {
		db.Exec(
			SQL_UPDATE,
			article_id,
			args.Peek("title"),
			body,
			body_ht,
			args.Has("public"),
		)
	}

	// Tags

	db.Exec(`DELETE FROM "tag"  WHERE "article_id" = $1`, article_id)

	for _, tag := range bytes.Split(args.Peek("tags"), []byte(",")) {
		tag = bytes.TrimSpace(tag)

		if len(tag) == 0 {
			continue
		}

		db.Exec(
			`INSERT INTO "tag"  ("article_id", "tag")  VALUES ($1, $2)  ON CONFLICT  DO NOTHING`,
			article_id, tag,
		)
	}

	db.Commit()

	return EditGet(ctx)
}
