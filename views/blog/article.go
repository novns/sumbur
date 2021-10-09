package blog

import (
	"sumbur/db"
	"time"

	"github.com/MakeNowJust/heredoc"
)

var SQL_ARTICLE = heredoc.Doc(`
SELECT
	"article"."article_id",
	"article"."title",
	"article"."body",
	"article"."body_ht",
	"article"."public",
	"article"."created",
	array_agg("tag"."tag"  ORDER BY "tag"."tag")  AS "tags"
FROM "article"
FULL OUTER JOIN "tag"  USING ("article_id")
WHERE
	"article"."article_id" = $1
GROUP BY
	"article"."article_id"
`)

type Article struct {
	stag *string

	article_id int
	title      []byte
	body       []byte
	body_ht    []byte
	public     bool
	created    time.Time
	tags       [][]byte
}

func (article *Article) QueryArticle(db *db.DB, article_id int) bool {
	return db.
		Query(SQL_ARTICLE, article_id).
		Get(
			&article.article_id,
			&article.title,
			&article.body,
			&article.body_ht,
			&article.public,
			&article.created,
			&article.tags,
		)
}
