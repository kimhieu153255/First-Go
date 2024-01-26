-- name: CreateUser :one
INSERT INTO users (email, full_name, password) VALUES ($1, $2, $3) RETURNING *;

-- name: SelectUser :one
select * from users where email = $1 limit 1;

-- name: GetListUsers :many
select * from users order by id;

-- name: DeleteUser :exec
delete from users where email = $1;

-- name: UpdateUser :one
update users set full_name = $2, password = $3 where email = $1 returning *;


-- --TO UPDATE NO DEADLOCK:
-- SELECT * FROM users WHERE email = $1 FOR NO KEY UPDATE;