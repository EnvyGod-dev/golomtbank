-- name: CreateUser :one
INSERT INTO "User" (
    "Username",
    "HashedPassword",
    "FirstName",
    "LastName",
    "Email"
)
VALUES (
    sqlc.arg('Username'),
    sqlc.arg('HashedPassword'),
    sqlc.arg('FirstName'),
    sqlc.arg('LastName'),
    sqlc.arg('Email')
) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM 
    "User"
WHERE
    "Username" = sqlc.arg('Username')
LIMIT 1;