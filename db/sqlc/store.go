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
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// txName := ctx.Value(txKey)
		// fmt.Println(txName, "create transfer")

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		// fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(AddMoneyParams{
				ctx:        ctx,
				q:          q,
				accountID1: arg.FromAccountID,
				amount1:    -arg.Amount,
				accountID2: arg.ToAccountID,
				amount2:    arg.Amount,
			})
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(AddMoneyParams{
				ctx:        ctx,
				q:          q,
				accountID1: arg.ToAccountID,
				amount1:    arg.Amount,
				accountID2: arg.FromAccountID,
				amount2:    -arg.Amount,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return result, err
}

type AddMoneyParams struct {
	ctx        context.Context
	q          *Queries
	accountID1 int64
	amount1    int64
	accountID2 int64
	amount2    int64
}

func addMoney(arg AddMoneyParams) (account1 Account, account2 Account, err error) {
	account1, err = arg.q.AddAccountBalance(arg.ctx, AddAccountBalanceParams{
		ID:     arg.accountID1,
		Amount: arg.amount1,
	})

	if err != nil {
		return
	}

	account2, err = arg.q.AddAccountBalance(arg.ctx, AddAccountBalanceParams{
		ID:     arg.accountID2,
		Amount: arg.amount2,
	})

	if err != nil {
		return
	}

	return
}
