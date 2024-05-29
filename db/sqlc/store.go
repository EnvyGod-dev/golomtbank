package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64  `json:"from_account_id"`
	ToAccountId   int64  `json:"to_account_id"`
	BankName      string `json:"bank_name"`
	Amount        int64  `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
	FeeEntry    Entry    `json:"fee_entry,omitempty"`
}

// TransferTx performs a money transfer from one account to another
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.Amount <= 0 {
			return fmt.Errorf("Үлдэгдэл 0-ээс их байх ёстой")
		}

		fromAccount, err := q.GetAccountById(ctx, arg.FromAccountId)
		if err != nil {
			return err
		}

		if fromAccount.Balance < arg.Amount {
			return fmt.Errorf("Үлдэгдэл хүрэлцэхгүй байна")
		}

		var fee int64
		if arg.BankName != "Голомт Банк" {
			fee = 100
		}

		result.Transfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountId: arg.FromAccountId,
			ToAccountId:   arg.ToAccountId,
			BankName:      arg.BankName,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			FromAccountId: arg.FromAccountId,
			ToAccountId:   arg.ToAccountId,
			BankName:      arg.BankName,
			Amount:        -(arg.Amount + fee),
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			FromAccountId: arg.ToAccountId,
			ToAccountId:   arg.FromAccountId,
			BankName:      arg.BankName,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		if fee > 0 {
			result.FeeEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
				FromAccountId: arg.FromAccountId,
				ToAccountId:   arg.ToAccountId,
				BankName:      arg.BankName,
				Amount:        -fee,
			})
			if err != nil {
				return err
			}
		}

		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountId, -(arg.Amount + fee), arg.ToAccountId, arg.Amount)
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountId, arg.Amount, arg.FromAccountId, -(arg.Amount + fee))
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountId1 int64, amount1 int64, accountId2 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Id:     accountId1,
		Amount: amount1,
	})
	if err != nil {
		return
	}
	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Id:     accountId2,
		Amount: amount2,
	})
	return
}
