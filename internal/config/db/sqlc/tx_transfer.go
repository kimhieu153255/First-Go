package db

import (
	"context"
)

type TransferTxResult struct {
	FromAccount Account
	ToAccount   Account
	Transfer    Transfer
	FromEntry   Entry
	ToEntry     Entry
}

func (store *SQLStore) TransferTx(ctx context.Context, arg AddTransferParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.AddTransfer(ctx, arg)

		if err != nil {
			return err
		}

		// create a new entry
		result.FromEntry, err = q.AddEntry(ctx, AddEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
			Currency:  arg.Currency,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.AddEntry(ctx, AddEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
			Currency:  arg.Currency,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = changeAccountBalance(ctx, q, arg.FromAccountID, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = changeAccountBalance(ctx, q, arg.ToAccountID, arg.FromAccountID, arg.Amount)
		}
		return err
	})
	return result, err
}

func changeAccountBalance(ctx context.Context, q *Queries, fromAccountId, toAccountID, amount int64) (Account, Account, error) {
	fromAccount, err := q.AddAccountBalanceByID(ctx, AddAccountBalanceByIDParams{
		Amount: -amount,
		ID:     fromAccountId,
	})
	if err != nil {
		return fromAccount, Account{}, err
	}

	toAccount, err := q.AddAccountBalanceByID(ctx, AddAccountBalanceByIDParams{
		Amount: amount,
		ID:     toAccountID,
	})
	return fromAccount, toAccount, err
}
