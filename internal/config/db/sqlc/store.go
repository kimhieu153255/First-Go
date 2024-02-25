package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	UpdateUserUseStore(ctx context.Context, arg UpdateUserTxParams) (UpdateUserResult, error)
	TransferTx(ctx context.Context, arg AddTransferParams) (TransferTxResult, error)
}

// Store provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewStore creates a new Store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  New(connPool),
		connPool: connPool,
	}
}

type UpdateUserTxParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type UpdateUserResult struct {
	User User `json:"user"`
}

func (store *SQLStore) UpdateUserUseStore(ctx context.Context, arg UpdateUserTxParams) (UpdateUserResult, error) {
	var result UpdateUserResult
	// to start a transaction we need to use ExecTx (in this, we have a lot of queries)
	err := store.ExecTx(ctx, func(q *Queries) error {

		_, err := q.SelectUserForUpdate(ctx, arg.ID)
		if err != nil {
			return err
		}

		user, err := q.UpdateUser(ctx, UpdateUserParams{
			ID:       arg.ID,
			Password: arg.Password,
			FullName: arg.FullName,
		})

		if err != nil {
			return err
		}

		fmt.Println("user", user)
		result.User = user
		return nil
	})

	return result, err
}
