CREATE TABLE `usuario` (
  `nome_usuario` varchar(255) PRIMARY KEY,
  `senha_hash` varchar(255) NOT NULL,
  `nome_completo` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `mudanca_senha` timestamptz NOT NULL DEFAULT "0001-01-01 00:00:00Z",
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `contas` (
  `id` bigserial PRIMARY KEY,
  `dono` varchar(255) NOT NULL,
  `saldo` bigint NOT NULL,
  `moeda` varchar(255) NOT NULL,
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `mudancas` (
  `id` bigserial PRIMARY KEY,
  `id_conta` bigint NOT NULL,
  `quantia` bigint NOT NULL COMMENT 'esse valor pode ser negativo ou positivo',
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `transferencias` (
  `id` bigserial PRIMARY KEY,
  `de_id_conta` bigint NOT NULL,
  `para_id_conta` bigint NOT NULL,
  `quantia` bigint NOT NULL COMMENT 'esse valor so pode ser positivo',
  `criada_em` timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE `contas` ADD FOREIGN KEY (`dono`) REFERENCES `usuario` (`nome_usuario`);

ALTER TABLE `mudancas` ADD FOREIGN KEY (`id_conta`) REFERENCES `contas` (`id`);

ALTER TABLE `transferencias` ADD FOREIGN KEY (`de_id_conta`) REFERENCES `contas` (`id`);

ALTER TABLE `transferencias` ADD FOREIGN KEY (`para_id_conta`) REFERENCES `contas` (`id`);

CREATE INDEX `contas_index_0` ON `contas` (`dono`);

CREATE UNIQUE INDEX `contas_index_1` ON `contas` (`dono`, `moeda`);

CREATE INDEX `mudancas_index_2` ON `mudancas` (`id_conta`);

CREATE INDEX `transferencias_index_3` ON `transferencias` (`de_id_conta`);

CREATE INDEX `transferencias_index_4` ON `transferencias` (`para_id_conta`);

CREATE INDEX `transferencias_index_5` ON `transferencias` (`de_id_conta`, `para_id_conta`);
