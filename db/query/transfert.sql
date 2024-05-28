-- name: CreateTransfers :one

INSERT INTO  "transfer"
    (
        "FromAccountId",
        "ToAccountId",
        "Amount"
    )
VALUES 
    (
        sqlc.arg('FromAccountId'),
        sqlc.arg('ToAccountId'),
        sqlc.arg('Amount')
    ) RETURNING *;

-- name: GetTransfers :one

SELECT 
    *
FROM 
    "transfer"
WHERE
    "Id" = sqlc.arg('Id')
LIMIT 1;

-- name: ListTransfers :many

SELECT
    *
FROM
    "transfer"
ORDER BY
    "Id" = sqlc.arg('Id') DESC
LIMIT 1
OFFSET 2;