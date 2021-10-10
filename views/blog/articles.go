package blog

import (
	_ "embed"
	"sumbur/db"
	"sumbur/views"
)

//go:embed sql/articles.sql
var SQL_ARTICLES string

//go:embed sql/articles-tag.sql
var SQL_ARTICLES_TAG string

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
