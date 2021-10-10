INSERT INTO "article" (
    "title",
    "body",
    "body_ht",
    "public",
    "created"
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
ON CONFLICT  DO NOTHING
RETURNING "article_id"
