package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/token"
)

func adcAutorizacao(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	nome_usuario string,
	duracao time.Duration,
) {
	token, payload, err := tokenMaker.CriarToken(nome_usuario, duracao)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testarCasos := []struct {
		nome           string
		setupAuth      func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checarResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			nome: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, "usuario", time.Minute)
			},
			checarResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			nome: "SemAutorizacao",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			checarResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			nome: "AutorizacaoSemSuporte",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, "semsuporte", "usuario", time.Minute)
			},
			checarResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			nome: "FormatoDaAutorizacaoInvalido",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, "", "usuario", time.Minute)
			},
			checarResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			nome: "TokenExpirado",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				adcAutorizacao(t, request, tokenMaker, authorizationTypeBearer, "usuario", -time.Minute)
			},
			checarResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testarCasos {
		tc := testarCasos[i]

		t.Run(tc.nome, func(t *testing.T) {
			servidor := newTestServer(t, nil)

			authPath := "/auth"
			servidor.roteador.GET(
				authPath,
				authMiddleware(servidor.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, servidor.tokenMaker)
			servidor.roteador.ServeHTTP(recorder, request)
			tc.checarResponse(t, recorder)
		})
	}
}
