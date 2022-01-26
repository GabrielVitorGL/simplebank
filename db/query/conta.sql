-- name: CriarConta :one
INSERT INTO contas (
  dono,
  saldo,
  moeda
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ObterConta :one
SELECT * FROM contas
WHERE id = $1 
LIMIT 1;

-- name: ListarContas :many
SELECT * FROM contas
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AtualizarConta :one
UPDATE contas 
SET saldo = $2
WHERE id = $1
RETURNING *;

-- name: DeletarConta :exec
DELETE FROM contas 
WHERE id = $1;