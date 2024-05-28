// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccounts(ctx context.Context, id int64) error
	GetAccountById(ctx context.Context, id int64) (Account, error)
	GetAccountsForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfers(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, id int64) ([]Entry, error)
	ListTransfers(ctx context.Context, id int64) ([]Transfer, error)
	UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Account, error)
}

var _ Querier = (*Queries)(nil)