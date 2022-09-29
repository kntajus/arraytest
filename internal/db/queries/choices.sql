-- name: AddChoice :exec
INSERT INTO choices (fruits)
VALUES ($1);
