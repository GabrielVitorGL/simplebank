package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/techschool/simplebank/db/mock"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/util"
)

func TestContaObtidaAPI(t *testing.T) {
	usuario, _ := usuarioAleatorio(t)
	conta := contaAleatoria(usuario.NomeUsuario)

	testarCasos := []struct {
		nome          string
		IDconta       int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			nome:    "OK",
			IDconta: conta.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, usuario.NomeUsuario, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ObterConta(gomock.Any(), gomock.Eq(conta.ID)). // obtendo uma conta com um contexto aleatório, e com o id da conta gerada
					Times(1).
					Return(conta, nil) // deve conter os argumentos de retorno do tipo "ObterConta" do arquivo querier.go
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//checar se deu certo
				require.Equal(t, http.StatusOK, recorder.Code)
				requerirConteudoContasIguais(t, recorder.Body, conta)
			},
		},
		{
			nome:    "UsuarioSemAutorizacao",
			IDconta: conta.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, "sem_autorizacao", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ObterConta(gomock.Any(), gomock.Eq(conta.ID)). // obtendo uma conta com um contexto aleatório, e com o id da conta gerada
					Times(1).
					Return(conta, nil) // deve conter os argumentos de retorno do tipo "ObterConta" do arquivo querier.go
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//checar se deu certo
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			nome:    "SemAutorizacao",
			IDconta: conta.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ObterConta(gomock.Any(), gomock.Any()). // obtendo uma conta com um contexto aleatório, e com o id da conta gerada
					Times(0) 
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//checar se deu certo
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			nome:    "NaoEncontrada",
			IDconta: conta.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, usuario.NomeUsuario, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ObterConta(gomock.Any(), gomock.Eq(conta.ID)). // obtendo uma conta com um contexto aleatório, e com o id da conta gerada
					Times(1).
					Return(db.Conta{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//checar se deu certo
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			nome:    "ErroInterno",
			IDconta: conta.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, usuario.NomeUsuario, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ObterConta(gomock.Any(), gomock.Eq(conta.ID)). // obtendo uma conta com um contexto aleatório, e com o id da conta gerada
					Times(1).
					Return(db.Conta{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//checar se deu certo
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			nome:    "IDInvalido",
			IDconta: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, usuario.NomeUsuario, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ObterConta(gomock.Any(), gomock.Any()). // obtendo uma conta com um contexto aleatório, e com o id da conta gerada
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//checar se deu certo
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testarCasos {
		tc := testarCasos[i]

		t.Run(tc.nome, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Iniciar o teste do servidor e mandar um request
			servidor := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/contas/%d", tc.IDconta)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, servidor.tokenMaker)
			servidor.roteador.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func contaAleatoria(dono string) db.Conta {
	return db.Conta{
		ID:    util.RandomInt(1, 1000), // pegando um id aleatorio para conta de 1 a 1000
		Dono:  dono,
		Saldo: util.RandomMoney(),
		Moeda: util.RandomCurrency(),
	}
}

func requerirConteudoContasIguais(t *testing.T, body *bytes.Buffer, conta db.Conta) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var contaObtida db.Conta
	err = json.Unmarshal(data, &contaObtida)
	require.NoError(t, err)

	require.Equal(t, conta, contaObtida)
}
