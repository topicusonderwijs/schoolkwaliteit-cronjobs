-- name: GetExpiredAccounts :many
SELECT accounts.*
FROM accounts
WHERE deleted_at IS NULL AND accounts.active = TRUE AND (expired_on < NOW());

-- name: DeactivateExpiredAccountById :exec
UPDATE accounts
SET active = false
WHERE id = ?
