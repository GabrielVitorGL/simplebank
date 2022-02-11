-- name: CriarUsuario :one
INSERT INTO usuario (
  nome_usuario,
  senha_hash,
  nome_completo,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ObterUsuario :one
SELECT * FROM usuario
WHERE nome_usuario = $1 
LIMIT 1;