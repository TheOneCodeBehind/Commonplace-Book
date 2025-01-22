-- name: CreatePosts :one
INSERT INTO posts (
    id, post, author, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: ListPostsByAuthor :many
SELECT * FROM posts
WHERE author = $1;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: UpdatePostById :one
UPDATE posts
    set post = $2
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

