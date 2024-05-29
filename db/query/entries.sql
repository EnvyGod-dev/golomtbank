-- name: CreateEntries :one
INSERT INTO
    "entries" (
        "FromAccountId",
        "BankName",
        "ToAccountId",
        "Amount"
    )
VALUES
    (
        sqlc.arg('FromAccountId'),
        sqlc.arg('BankName'),
        sqlc.arg('ToAccountId'),
        sqlc.arg('Amount')
    ) RETURNING *;

-- name: GetEntry :one
SELECT
    *
FROM
    "entries"
WHERE
    "Id" = sqlc.arg('Id')
LIMIT
    1;

-- name: ListEntries :many
SELECT
    *
FROM
    "entries"
ORDER BY
    "Id" = sqlc.arg('Id') DESC
LIMIT
    1 OFFSET 2;