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
