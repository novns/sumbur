package blog

import (
	_ "embed"
	"sumbur/db"
	"time"
)

//go:embed sql/article.sql
var SQL_ARTICLE string

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
