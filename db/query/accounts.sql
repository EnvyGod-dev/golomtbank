-- name: CreateAccounts :one
INSERT INTO
    "accounts" (
        "Id",
        "Balance",
        "Owner",
        "BankName",
        "Currency"
    )
VALUES
    (
        sqlc.arg('Id'),
        sqlc.arg('Balance'),
        sqlc.arg('Owner'),
        sqlc.arg('BankName'),
        sqlc.arg('Currency')
    ) RETURNING *;

-- name: GetAccountById :one
SELECT
    *
FROM
    accounts
WHERE
    "Id" = $1
LIMIT
    1;

-- name: UpdateAccounts :one
UPDATE
    "accounts"
SET
    "Balance" = sqlc.arg('Balance')
WHERE
    "Id" = sqlc.arg('Id') RETURNING *;

-- name: ListAccounts :many
SELECT
    *
FROM
    accounts
WHERE
    "Owner" = sqlc.arg('Owner')
ORDER BY
    "Id"
LIMIT
    $1
OFFSET
    $2;

-- name: DeleteAccounts :exec
DELETE FROM
    "accounts"
WHERE
    "Id" = sqlc.arg('Id');

-- name: GetAccountsForUpdate :one
SELECT
    *
FROM
    "accounts"
WHERE
    "Id" = sqlc.arg('Id')
LIMIT
    1 FOR NO KEY
UPDATE
;

-- name: AddAccountBalance :one
UPDATE
    accounts
SET
    "Balance" = "Balance" + sqlc.arg('amount')
WHERE
    "Id" = sqlc.arg('Id') RETURNING *;