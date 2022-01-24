CREATE TABLE "contas" (
  "id" bigserial PRIMARY KEY,
  "dono" varchar NOT NULL,
  "saldo" bigint NOT NULL,
  "moeda" varchar NOT NULL,
  "criada_em" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "mudancas" (
  "id" bigserial PRIMARY KEY,
  "id_conta" bigint NOT NULL,
  "quantia" bigint NOT NULL,
  "criada_em" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transferencias" (
  "id" bigserial PRIMARY KEY,
  "de_id_conta" bigint NOT NULL,
  "para_id_conta" bigint NOT NULL,
  "quantia" bigint NOT NULL,
  "criada_em" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "mudancas" ADD FOREIGN KEY ("id_conta") REFERENCES "contas" ("id");

ALTER TABLE "transferencias" ADD FOREIGN KEY ("de_id_conta") REFERENCES "contas" ("id");

ALTER TABLE "transferencias" ADD FOREIGN KEY ("para_id_conta") REFERENCES "contas" ("id");

CREATE INDEX ON "contas" ("dono");

CREATE INDEX ON "mudancas" ("id_conta");

CREATE INDEX ON "transferencias" ("de_id_conta");

CREATE INDEX ON "transferencias" ("para_id_conta");

CREATE INDEX ON "transferencias" ("de_id_conta", "para_id_conta");

COMMENT ON COLUMN "mudancas"."quantia" IS 'esse valor pode ser negativo ou positivo';

COMMENT ON COLUMN "transferencias"."quantia" IS 'esse valor so pode ser positivo';
