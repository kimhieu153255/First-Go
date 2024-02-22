package db

import "context"

type TransferTxResult struct {
	FromAccount Account
	ToAccount   Account
	Transfer    Transfer
	FromEntry   Entry
	ToEntry     Entry
}

type TransferTxParams struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.AddTransfer(ctx, AddTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
			Currency:      arg.Currency,
		})
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

		result.FromAccount, err = q.AddAccountBalanceByID(ctx, AddAccountBalanceByIDParams{
			Amount: -arg.Amount,
			ID:     arg.FromAccountID,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalanceByID(ctx, AddAccountBalanceByIDParams{
			Amount: arg.Amount,
			ID:     arg.ToAccountID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
