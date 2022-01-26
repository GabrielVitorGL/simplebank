# simplebank
Simplebank project using Golang

## OBJETIVOS DESTE PROJETO:

1- Criar a conta da pessoa com seus dados, e armazenar seu saldo

2- Gravar todas as mudanças na conta, fazendo um registro a cada mudança que ocorrer

3- Transferência de dinheiro entre duas contas de forma funcional

## CONTEÚDO DAS AULAS JÁ REALIZADAS:

### 24/01

 #### Primeira aula - Database Design (Banco de dados)
   * Desenhar um esquema de Banco de Dados do SQL no dbdiagram.io
   * Salvar esse esquema como PDF ou PNG e compartilhá-lo
   * Gerar o código desse esquema no SQL Server

 #### Segunda aula - Docker + Postgres + TablePlus
  * Instalar o Docker 
  * Instalar e iniciar um container do PostgreSQL
  * Utilizar alguns comandos básicos do docker
  * Preparar e utilizar o TablePlus pra conectar e interagir com o Postgres
  * Usar o TablePlus para rodar o script SQL que geramos na aula 1 para criar o esquema de banco de dados do nosso projeto

 #### Terceira aula - Escrevendo e rodando a migração do banco de dados
  * Instalar a biblioteca Migrate
  * Escrever os códigos que iremos utilizar
  * Rodar os comandos
  * Criar um makefile para rodar todos os comandos automaticamente usando atalhos sempre que precisarmos iniciar ou fechar o nosso banco de dados
 
###25/01
 
 ####Quarta aula - Fazendo um código em Go para realizar o CRUD do nosso banco de dados
  *Aprender como utilizar a biblioteca SQLC para manusear o banco de dados
  *Fazer o código que realizará a inserção, leitura atualização e exclusão de informações no nosso banco de dados
  *Rodar e testar o código
  
 #### Quinta aula - Inserindo valores no banco de dados usando o código da última aula
  * Criar o código do arquivo main_test
  * Criar o código do arquivo conta_test
  * Testar os códigos
  * Fazer um código que gerará diversos dados aleatórios
  * Escrever esses dados aleatórios no nosso banco de dados para checar se há erros
  * Obs: a partir de agora usando o DBeaver para a visualização do banco de dados
  
  ##### Exercício passado realizado depois da aula:
   * Criar o código do arquivo mudanca_test
   * Criar o código do arquivo transferencia_test
   * Testar todos os códigos, que dessa vez estão preenchendo todos os dados do nosso banco de dados com valores aleatórios para teste
