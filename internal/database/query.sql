-- name: GetPost :one 
SELECT title, content, description, created_at, updated_at
FROM posts
WHERE id = ?1 AND removed_at IS NULL;

-- name: ListPosts :many 
SELECT id, title, content, description, created_at, updated_at
FROM posts
WHERE removed_at IS NULL;

-- name: CreatePost :one
INSERT INTO posts
  (title, content, description) VALUES
  (?1, ?2, ?3)
RETURNING id;

-- name: UpdatePost :exec 
UPDATE posts
SET
  title = CASE
    WHEN CAST(@title AS TEXT) = '' THEN title
    ELSE CAST(@title AS TEXT)
  END,
  content = CASE
    WHEN CAST(@content AS TEXT) = '' THEN content
    ELSE CAST(@content AS TEXT)
  END,
  description = CASE
    WHEN CAST(@description AS TEXT) = '' THEN description
    ELSE CAST(@description AS TEXT)
  END,
  updated_at = CURRENT_TIMESTAMP
WHERE id = ?1 AND removed_at IS NULL;

-- name: DeletePost :exec
UPDATE posts
SET
  removed_at = CURRENT_TIMESTAMP
WHERE id = ?1 AND removed_at IS NULL;
