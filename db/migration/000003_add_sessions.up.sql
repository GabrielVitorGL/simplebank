CREATE TABLE "secoes" (
  "id" uuid PRIMARY KEY,
  "nome_usuario" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar UNIQUE NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT FALSE,
  "expira_em" timestamptz NOT NULL,
  "criada_em" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "secoes" ADD FOREIGN KEY ("nome_usuario") REFERENCES "usuario" ("nome_usuario");
