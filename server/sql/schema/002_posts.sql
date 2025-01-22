-- +goose Up
CREATE TABLE posts(
    "id" UUID PRIMARY KEY,
    "post" TEXT NOT NULL,
    "author" UUID NOT NULL REFERENCES users (id),
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE posts;