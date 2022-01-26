-- name: CriarMudanca :one
INSERT INTO mudancas (
  id_conta,
  quantia
) VALUES (
  $1, $2
) RETURNING *;

-- name: ObterMudanca :one
SELECT * FROM mudancas
WHERE id = $1 
LIMIT 1;

-- name: ListarMudancas :many
SELECT * FROM mudancas
WHERE id_conta = $1
ORDER BY id
LIMIT $2
OFFSET $3;