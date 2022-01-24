CREATE TABLE [contas] (
  [id] bigserial PRIMARY KEY,
  [dono] nvarchar(255) NOT NULL,
  [saldo] bigint NOT NULL,
  [moeda] nvarchar(255) NOT NULL,
  [criada_em] timestamptz NOT NULL DEFAULT (now())
)
GO

CREATE TABLE [mudancas] (
  [id] bigserial PRIMARY KEY,
  [id_conta] bigint,
  [quantia] bigint NOT NULL,
  [criada_em] timestamptz NOT NULL DEFAULT (now())
)
GO

CREATE TABLE [transferencias] (
  [id] bigserial PRIMARY KEY,
  [de_id_conta] bigint,
  [para_id_conta] bigint,
  [quantia] bigint NOT NULL,
  [criada_em] timestamptz NOT NULL DEFAULT (now())
)
GO

ALTER TABLE [mudancas] ADD FOREIGN KEY ([id_conta]) REFERENCES [contas] ([id])
GO

ALTER TABLE [transferencias] ADD FOREIGN KEY ([de_id_conta]) REFERENCES [contas] ([id])
GO

ALTER TABLE [transferencias] ADD FOREIGN KEY ([para_id_conta]) REFERENCES [contas] ([id])
GO

CREATE INDEX [contas_index_0] ON [contas] ("dono")
GO

CREATE INDEX [mudancas_index_1] ON [mudancas] ("id_conta")
GO

CREATE INDEX [transferencias_index_2] ON [transferencias] ("de_id_conta")
GO

CREATE INDEX [transferencias_index_3] ON [transferencias] ("para_id_conta")
GO

CREATE INDEX [transferencias_index_4] ON [transferencias] ("de_id_conta", "para_id_conta")
GO

EXEC sp_addextendedproperty
@name = N'Column_Description',
@value = 'esse valor pode ser negativo ou positivo',
@level0type = N'Schema', @level0name = 'dbo',
@level1type = N'Table',  @level1name = 'mudancas',
@level2type = N'Column', @level2name = 'quantia';
GO

EXEC sp_addextendedproperty
@name = N'Column_Description',
@value = 'esse valor so pode ser positivo',
@level0type = N'Schema', @level0name = 'dbo',
@level1type = N'Table',  @level1name = 'transferencias',
@level2type = N'Column', @level2name = 'quantia';
GO
