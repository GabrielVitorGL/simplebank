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

-- name: ObterContaParaAtualizar :one
SELECT * FROM contas
WHERE id = $1 
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListarContas :many
SELECT * FROM contas
WHERE dono = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: AtualizarConta :one
UPDATE contas 
SET saldo = $2
WHERE id = $1
RETURNING *;

-- name: AdicionarSaldoConta :one
UPDATE contas 
SET saldo = saldo + sqlc.arg(quantia)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletarConta :exec
DELETE FROM contas 
WHERE id = $1;