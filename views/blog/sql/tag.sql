SELECT "tag"."tag"
FROM "tag"
INNER JOIN "article"  USING ("article_id")
WHERE
    ($1 OR "article"."public") AND ("tag"."tag" = $2)
LIMIT 1
