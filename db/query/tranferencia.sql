-- name: CriarTransferencia :one
INSERT INTO transferencias (
  de_id_conta,
  para_id_conta,
  quantia
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ObterTransferencia :one
SELECT * FROM transferencias
WHERE id = $1 
LIMIT 1;

-- name: ListarTransferencias :many
SELECT * FROM transferencias
WHERE 
    de_id_conta = $1 OR
    para_id_conta = $2
ORDER BY id
LIMIT $3
OFFSET $4;