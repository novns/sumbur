CREATE TABLE "article" (
    "article_id"  serial  PRIMARY KEY,

    "title"  varchar  NOT NULL  DEFAULT '',
    "body"  text  NOT NULL  DEFAULT '',
    "body_ht"  text  NOT NULL  DEFAULT '',

    "public"  boolean  NOT NULL  DEFAULT FALSE,

    "created"  timestamptz  NOT NULL  DEFAULT now()
);

CREATE INDEX "article_public_idx"  ON "article" ("public");
CREATE INDEX "article_created_idx"  ON "article" ("created");


CREATE TABLE "tag" (
    "article_id"  integer  NOT NULL,
    "tag"  varchar  NOT NULL,

    PRIMARY KEY ("article_id", "tag"),

    CONSTRAINT "tag_article_id_fkey"
        FOREIGN KEY ("article_id")  REFERENCES "article" ("article_id")
        ON DELETE CASCADE
);

CREATE INDEX "tag_article_id"  ON "tag" ("article_id");
CREATE INDEX "tag_tag"  ON "tag" ("tag");

--

INSERT INTO "article" (
    "article_id",
    "title",
    "body",
    "body_ht",
    "public",
    "created"
)
VALUES
(
    1,
    'Title 1',
    'Article 1',
    '<p>Article 1</p>',
    't',
    '2021-09-30 00:01:00'
),
(
    2,
    'Title 2',
    'Article 2',
    '<p>Article 2</p>',
    't',
    '2021-09-30 00:02:00'
),
(
    3,
    'Title 3',
    'Article 3',
    '<p>Article 3</p>',
    'f',
    '2021-09-30 00:03:00'
);

INSERT INTO "tag" ("article_id", "tag")
VALUES
    (1, 'Tag 1'),
    (1, 'Tag 2'),
    (2, 'Tag 2'),
    (2, 'Tag 3'),
    (3, 'Tag 3'),
    (3, 'Tag 4')
;
