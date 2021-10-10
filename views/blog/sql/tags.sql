SELECT
    "tag",
    count("article_id")  AS "cnt"
FROM "tag"
INNER JOIN "article"  USING ("article_id")
WHERE
    $1 OR "article"."public"
GROUP BY
    "tag"
ORDER BY
    "cnt" DESC,
    "tag"
