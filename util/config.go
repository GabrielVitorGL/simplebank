package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config irá guardar todas as configurações da aplicação
// Os valores serão lidos pelo viper de um arquivo config ou uma variável de ambiente
type Config struct {
	DBDriver           string        `mapstructure:"DB_DRIVER"`
	DBSource           string        `mapstructure:"DB_SOURCE"`
	EnderecoServidor   string        `mapstructure:"ENDERECO_SERVIDOR"`
	TokenSymmetricKey  string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACESS_TOKEN_DURATION"`
}

func CarregarConfig(caminho string) (config Config, err error) {
	viper.AddConfigPath(caminho)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // pode ser json, xml etc

	viper.AutomaticEnv() // Substituir os valores do arquivo config por o valor das variáveis de ambiente, se eles existirem

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
