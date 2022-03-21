-- name: CreateSession :one
INSERT INTO secoes (
  id,
  nome_usuario,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expira_em
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM secoes
WHERE id = $1 LIMIT 1;
