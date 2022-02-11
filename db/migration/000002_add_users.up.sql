CREATE TABLE "usuario" (
  "nome_usuario" varchar PRIMARY KEY,
  "senha_hash" varchar NOT NULL,
  "nome_completo" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "mudanca_senha" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "criada_em" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "contas" ADD FOREIGN KEY ("dono") REFERENCES "usuario" ("nome_usuario");

CREATE UNIQUE INDEX ON "contas" ("dono", "moeda");