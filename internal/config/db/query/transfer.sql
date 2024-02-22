-- name: AddTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount, currency)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransferByID :one
select * from transfers where id = $1 limit 1;

-- name: GetTransfersByAccountID :many
select * from transfers
where from_account_id = $1 or to_account_id = $1
order by id;

-- name: GetTransfers :many
select * from transfers order by id;

-- name: DeleteTransferByID :one
delete from transfers where id = $1 returning *;