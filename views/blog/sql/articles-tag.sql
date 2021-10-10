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
