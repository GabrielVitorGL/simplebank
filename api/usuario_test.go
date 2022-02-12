package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	mockdb "github.com/techschool/simplebank/db/mock"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

type eqCriarUsuarioParamsMatcher struct {
	arg   db.CriarUsuarioParams
	senha string
}

func (e eqCriarUsuarioParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CriarUsuarioParams)
	if !ok {
		return false
	}

	err := util.ChecarSenha(e.senha, arg.SenhaHash)
	if err != nil {
		return false
	}

	e.arg.SenhaHash = arg.SenhaHash
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCriarUsuarioParamsMatcher) String() string {
	return fmt.Sprintf("matches argumento %v e senha %v", e.arg, e.senha)
}

func EqCriarUsuarioParams(arg db.CriarUsuarioParams, senha string) gomock.Matcher {
	return eqCriarUsuarioParamsMatcher{arg, senha}
}

func TestCriarUsuarioAPI(t *testing.T) {
	usuario, senha := usuarioAleatorio(t)

	testarCasos := []struct {
		nome           string
		conteudo       gin.H
		buildStubs     func(store *mockdb.MockStore)
		checarResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			nome: "OK",
			conteudo: gin.H{
				"nome_usuario":  usuario.NomeUsuario,
				"senha":         senha,
				"nome_completo": usuario.NomeCompleto,
				"email":         usuario.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CriarUsuarioParams{
					NomeUsuario:  usuario.NomeUsuario,
					NomeCompleto: usuario.NomeCompleto,
					Email:        usuario.Email,
				}
				store.EXPECT().
					CriarUsuario(gomock.Any(), EqCriarUsuarioParams(arg, senha)).
					Times(1).
					Return(usuario, nil)
			},
			checarResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, usuario)
			},
		},
		{
			nome: "InternalError",
			conteudo: gin.H{
				"nome_usuario":  usuario.NomeUsuario,
				"senha":         senha,
				"nome_completo": usuario.NomeCompleto,
				"email":         usuario.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CriarUsuario(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Usuario{}, sql.ErrConnDone)
			},
			checarResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			nome: "DuplicateUsername",
			conteudo: gin.H{
				"nome_usuario":  usuario.NomeUsuario,
				"senha":         senha,
				"nome_completo": usuario.NomeCompleto,
				"email":         usuario.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CriarUsuario(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Usuario{}, &pq.Error{Code: "23505"})
			},
			checarResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			nome: "InvalidUsername",
			conteudo: gin.H{
				"nome_usuario":  "usuario-invalido#1",
				"senha":         senha,
				"nome_completo": usuario.NomeCompleto,
				"email":         usuario.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CriarUsuario(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checarResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			nome: "InvalidEmail",
			conteudo: gin.H{
				"nome_usuario":  usuario.NomeUsuario,
				"senha":         senha,
				"nome_completo": usuario.NomeCompleto,
				"email":         "email-invalido",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CriarUsuario(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checarResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			nome: "TooShortPassword",
			conteudo: gin.H{
				"nome_usuario":  usuario.NomeUsuario,
				"senha":         "123",
				"nome_completo": usuario.NomeCompleto,
				"email":         usuario.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CriarUsuario(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checarResponse: func(recorder *httptest.ResponseRecorder) {
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
			servidor := NovoServidor(store)
			recorder := httptest.NewRecorder()
			
			data, err := json.Marshal(tc.conteudo)
			require.NoError(t, err)
			url := "/usuarios"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			servidor.roteador.ServeHTTP(recorder, request)
			tc.checarResponse(recorder)
		})
	}
}
func usuarioAleatorio(t *testing.T) (usuario db.Usuario, senha string) {
	n := util.RandomInt(6, 12)
	senha = util.RandomString(int(n))
	senhaHash, err := util.SenhaHash(senha)
	require.NoError(t, err)
	usuario = db.Usuario{
		NomeUsuario:       util.RandomOwner(),
		SenhaHash: senhaHash,
		NomeCompleto:       util.RandomOwner() + " " + util.RandomOwner(),
		Email:          util.EmailAleatorio(),
	}
	return
}
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, usuario db.Usuario) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var usuarioObtido db.Usuario
	err = json.Unmarshal(data, &usuarioObtido)
	require.NoError(t, err)
	require.Equal(t, usuario.NomeUsuario, usuarioObtido.NomeUsuario)
	require.Equal(t, usuario.NomeCompleto, usuarioObtido.NomeCompleto)
	require.Equal(t, usuario.Email, usuarioObtido.Email)
	require.Empty(t, usuarioObtido.SenhaHash)
}
