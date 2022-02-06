package db

import (
	"context"
	"database/sql"
	"fmt"
)
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}
// Store irá providenciar todas as funções para executar as queries do banco de dados e as transações
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// Executará uma função com uma transação do banco de dados
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil { // Se der um erro o banco de dados deve descartar as mudanças
			return fmt.Errorf("rx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	// Se todas as operações forem concluídas com sucesso:
	return tx.Commit() // Escrever as mudanças no banco de dados
}

// Contém os parâmetros para realizar uma transação
type TransferTxParams struct {
	DeIDConta   int64 `json:"de_id_conta"`
	ParaIDConta int64 `json:"para_id_conta"`
	Quantia     int64 `json:"quantia"`
}

// TransferTxResult é o resultado de uma transação
type TransferTxResult struct {
	Transferencia Transferencia `json:"transferencia"`
	DeConta       Conta         `json:"de_conta"`
	ParaConta     Conta         `json:"para_conta"`
	DeMudanca     Mudanca       `json:"de_mudanca"`
	ParaMudanca   Mudanca       `json:"para_mudanca"`
}

// TransferTx fará uma transação de valores de uma conta para a outra
// Criará um registro da transação e atualizará o valor na conta com uma única transação no banco de dados
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var resultado TransferTxResult // Criando uma variável do tipo TransferTxResult, por enquanto vazia

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		resultado.Transferencia, err = q.CriarTransferencia(ctx, CriarTransferenciaParams{
			DeIDConta:   arg.DeIDConta,
			ParaIDConta: arg.ParaIDConta,
			Quantia:     arg.Quantia,
		})
		if err != nil {
			return err
		}

		resultado.DeMudanca, err = q.CriarMudanca(ctx, CriarMudancaParams{
			IDConta: arg.DeIDConta,
			Quantia: -arg.Quantia,
		})
		if err != nil {
			return err
		}

		resultado.ParaMudanca, err = q.CriarMudanca(ctx, CriarMudancaParams{
			IDConta: arg.ParaIDConta,
			Quantia: arg.Quantia,
		})
		if err != nil {
			return err
		}

		// atualizar saldo
		if arg.DeIDConta < arg.ParaIDConta {
			resultado.DeConta, resultado.ParaConta, err = AdicionarDinheiro(ctx, q, arg.DeIDConta, -arg.Quantia, arg.ParaIDConta, arg.Quantia)
		} else {
			resultado.ParaConta, resultado.DeConta, err = AdicionarDinheiro(ctx, q, arg.ParaIDConta, arg.Quantia, arg.DeIDConta, -arg.Quantia)
		}

		return nil
	})

	return resultado, err
}

func AdicionarDinheiro (
	ctx context.Context,
	q *Queries,
	IDconta1 int64,
	quantia1 int64,
	IDconta2 int64,
	quantia2 int64,
 ) (conta1 Conta, conta2 Conta, err error) {
	conta1, err = q.AdicionarSaldoConta(ctx, AdicionarSaldoContaParams{
		ID: IDconta1,
		Quantia: quantia1,
	})
	if err != nil {
		return
	}

	conta2, err = q.AdicionarSaldoConta(ctx, AdicionarSaldoContaParams{
		ID: IDconta2,
		Quantia: quantia2,
	})
	return
}
