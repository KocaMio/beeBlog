-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE article (
    "id" serial not null primary key,
    "title" text not null,
    "content" text not null,
    "createTime" timestamp not null default CURRENT_TIMESTAMP,
    "editTime" timestamp not null default CURRENT_TIMESTAMP,
    "tag" json default null
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE article;
