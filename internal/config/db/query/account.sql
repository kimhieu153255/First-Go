-- name: CreateAccount :one
INSERT INTO accounts (user_id, balance, currency) VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccountByID :one
select * from accounts where id = $1 limit 1;

-- name: GetAccountByUserID :one
select * from accounts where user_id = $1 limit 1;

-- name: SelectAccountForUpdate :one
select * from accounts where id = $1 for no key update;

-- name: GetListAccounts :many
select * from accounts order by id;

-- name: DeleteAccountByID :one
delete from accounts where id = $1 returning *;

-- name: UpdateAccountBalanceByID :one
update accounts set balance = $2 where id = $1 returning *;

-- name: AddAccountBalanceByID :one
update accounts set balance = balance + sqlc.arg(amount) where id = sqlc.arg(id) returning *;