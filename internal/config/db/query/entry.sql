-- name: AddEntry :one
INSERT INTO entries (account_id, amount, currency)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetEntryByID :one
select * from entries where id = $1 limit 1;

-- name: GetEntriesByAccountID :many
select * from entries where account_id = $1 order by id;

-- name: GetEntries :many
select * from entries order by id;

-- name: DeleteEntryByID :one
delete from entries where id = $1 returning *;
