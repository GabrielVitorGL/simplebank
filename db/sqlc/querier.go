// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	AdicionarSaldoConta(ctx context.Context, arg AdicionarSaldoContaParams) (Conta, error)
	AtualizarConta(ctx context.Context, arg AtualizarContaParams) (Conta, error)
	CriarConta(ctx context.Context, arg CriarContaParams) (Conta, error)
	CriarMudanca(ctx context.Context, arg CriarMudancaParams) (Mudanca, error)
	CriarTransferencia(ctx context.Context, arg CriarTransferenciaParams) (Transferencia, error)
	DeletarConta(ctx context.Context, id int64) error
	ListarContas(ctx context.Context, arg ListarContasParams) ([]Conta, error)
	ListarMudancas(ctx context.Context, arg ListarMudancasParams) ([]Mudanca, error)
	ListarTransferencias(ctx context.Context, arg ListarTransferenciasParams) ([]Transferencia, error)
	ObterConta(ctx context.Context, id int64) (Conta, error)
	ObterContaParaAtualizar(ctx context.Context, id int64) (Conta, error)
	ObterMudanca(ctx context.Context, id int64) (Mudanca, error)
	ObterTransferencia(ctx context.Context, id int64) (Transferencia, error)
}

var _ Querier = (*Queries)(nil)
