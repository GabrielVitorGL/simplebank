CREATE TABLE `contas` (
  `id` bigserial PRIMARY KEY,
  `dono` varchar(255) NOT NULL,
  `saldo` bigint NOT NULL,
  `moeda` varchar(255) NOT NULL,
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `mudancas` (
  `id` bigserial PRIMARY KEY,
  `id_conta` bigint,
  `quantia` bigint NOT NULL COMMENT 'esse valor pode ser negativo ou positivo',
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `transferencias` (
  `id` bigserial PRIMARY KEY,
  `de_id_conta` bigint,
  `para_id_conta` bigint,
  `quantia` bigint NOT NULL COMMENT 'esse valor so pode ser positivo',
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE `mudancas` ADD FOREIGN KEY (`id_conta`) REFERENCES `contas` (`id`);

ALTER TABLE `transferencias` ADD FOREIGN KEY (`de_id_conta`) REFERENCES `contas` (`id`);

ALTER TABLE `transferencias` ADD FOREIGN KEY (`para_id_conta`) REFERENCES `contas` (`id`);

CREATE INDEX `contas_index_0` ON `contas` (`dono`);

CREATE INDEX `mudancas_index_1` ON `mudancas` (`id_conta`);

CREATE INDEX `transferencias_index_2` ON `transferencias` (`de_id_conta`);

CREATE INDEX `transferencias_index_3` ON `transferencias` (`para_id_conta`);

CREATE INDEX `transferencias_index_4` ON `transferencias` (`de_id_conta`, `para_id_conta`);
