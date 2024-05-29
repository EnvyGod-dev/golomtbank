// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: entries.sql

package db

import (
	"context"
)

const createEntries = `-- name: CreateEntries :one
INSERT INTO
    "entries" (
        "FromAccountId",
        "BankName",
        "Amount"
    )
VALUES
    (
        $1,
        $2,
        $3
    ) RETURNING "Id", "FromAccountId", "BankName", "Amount", "CreatedAt"
`

type CreateEntriesParams struct {
	FromAccountId int64  `json:"FromAccountId"`
	BankName      string `json:"BankName"`
	Amount        int64  `json:"Amount"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntries, arg.FromAccountId, arg.BankName, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.Id,
		&i.FromAccountId,
		&i.BankName,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
SELECT
    "Id", "FromAccountId", "BankName", "Amount", "CreatedAt"
FROM
    "entries"
WHERE
    "Id" = $1
LIMIT
    1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.Id,
		&i.FromAccountId,
		&i.BankName,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT
    "Id", "FromAccountId", "BankName", "Amount", "CreatedAt"
FROM
    "entries"
ORDER BY
    "Id" = $1 DESC
LIMIT
    1 OFFSET 2
`

func (q *Queries) ListEntries(ctx context.Context, id int64) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.Id,
			&i.FromAccountId,
			&i.BankName,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
