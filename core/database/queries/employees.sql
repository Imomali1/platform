-- name: DetailEmployee :one
SELECT * FROM EMPLOYEES
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM EMPLOYEES
ORDER BY employee_id;

-- name: CreateEmployee :exec
INSERT INTO EMPLOYEES
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- -- name: UpdateAuthor :exec
-- UPDATE EMPLOYEES
-- set name = $2,
--     bio = $3
-- WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM EMPLOYEES
WHERE id = $1;