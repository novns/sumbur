package blog

import (
	"sumbur/db"
	"sumbur/views"

	"github.com/MakeNowJust/heredoc"
)

var SQL_ARTICLES = heredoc.Doc(`
SELECT
	"article"."article_id",
	"article"."title",
	"article"."body_ht",
	"article"."public",
	"article"."created",
	array_agg("tag"."tag"  ORDER BY "tag"."tag")  AS "tags"
FROM "article"
FULL OUTER JOIN "tag"  USING ("article_id")
WHERE
	$1 OR "article"."public"
GROUP BY
	"article"."article_id"
ORDER BY
	"created" DESC
`)

var SQL_ARTICLES_TAG = heredoc.Doc(`
SELECT
	"article"."article_id",
	"article"."title",
	"article"."body_ht",
	"article"."public",
	"article"."created",
	array_agg("tag"."tag"  ORDER BY "tag"."tag")  AS "tags"
FROM "article"
FULL OUTER JOIN "tag"  USING ("article_id")
INNER JOIN "tag"  AS "tag_filter"  USING ("article_id")
WHERE
	($1 OR "article"."public") AND ("tag_filter"."tag" = $2)
GROUP BY
	"article"."article_id"
ORDER BY
	"created" DESC
`)

type Articles struct {
	query *db.Rows

	article Article
}

func QueryArticles(db *db.DB, stag *string) *Articles {
	articles := Articles{
		article: Article{stag: stag},
	}

	if stag == &views.EmptyString {
		articles.query = db.Query(SQL_ARTICLES, views.AuthState)
	} else {
		articles.query = db.Query(SQL_ARTICLES_TAG, views.AuthState, stag)
	}

	articles.query.Fields(
		&articles.article.article_id,
		&articles.article.title,
		&articles.article.body_ht,
		&articles.article.public,
		&articles.article.created,
		&articles.article.tags,
	)

	return &articles
}

func (articles *Articles) Next() bool {
	return articles.query.Next()
}
