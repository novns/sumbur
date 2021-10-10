UPDATE "article"
SET
    "title" = $2,
    "body" = $3,
    "body_ht" = $4,
    "public" = $5
WHERE "article_id" = $1
