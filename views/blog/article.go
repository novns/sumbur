package blog

import "time"

type Article struct {
	stag *string

	article_id int
	title      []byte
	body_ht    []byte
	public     bool
	created    time.Time
	tags       [][]byte
}
