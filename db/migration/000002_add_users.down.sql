DROP INDEX IF EXISTS contas_dono_moeda_idx;

ALTER TABLE IF EXISTS "contas" DROP CONSTRAINT IF EXISTS "contas_dono_fkey";

DROP TABLE IF EXISTS "usuario"