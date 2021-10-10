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
