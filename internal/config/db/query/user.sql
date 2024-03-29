-- name: CreateUser :one
INSERT INTO users (email, full_name, password) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserByEmail :one
select * from users where email = $1 limit 1;

-- name: GetUserById :one
select * from users where id = $1 limit 1;

-- name: SelectUserForUpdate :one
select * from users where id = $1 for no key update;

-- name: GetListUsers :many
select * from users order by id;

-- name: DeleteUserByID :one
delete from users where id = $1 returning *;

-- name: DeleteUserByEmail :one
delete from users where email = $1 returning *;

-- name: UpdateUser :one
update users set full_name = $2, password = $3 where id = $1 returning *;
