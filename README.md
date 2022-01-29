# simplebank
Simplebank project using Golang

## OBJETIVOS DESTE PROJETO:

1- Criar a conta da pessoa com seus dados, e armazenar seu saldo

2- Gravar todas as mudanças na conta, fazendo um registro a cada mudança que ocorrer

3- Transferência de dinheiro entre duas contas de forma funcional

## CONTEÚDO DAS AULAS JÁ REALIZADAS:

### 24/01

 #### 1º aula - Database Design (Banco de dados)
   * Desenhar um esquema de Banco de Dados do SQL no dbdiagram.io
   * Salvar esse esquema como PDF ou PNG e compartilhá-lo
   * Gerar o código desse esquema no SQL Server

 #### 2º aula - Docker + Postgres + TablePlus
  * Instalar o Docker 
  * Instalar e iniciar um container do PostgreSQL
  * Utilizar alguns comandos básicos do docker
  * Preparar e utilizar o TablePlus pra conectar e interagir com o Postgres
  * Usar o TablePlus para rodar o script SQL que geramos na aula 1 para criar o esquema de banco de dados do nosso projeto

 #### 3º aula - Escrevendo e rodando a migração do banco de dados
  * Instalar a biblioteca Migrate
  * Escrever os códigos que iremos utilizar
  * Rodar os comandos
  * Criar um makefile para rodar todos os comandos automaticamente usando atalhos sempre que precisarmos iniciar ou fechar o nosso banco de dados
 
### 25/01
 
 #### 4º aula - Fazendo um código em Go para realizar o CRUD do nosso banco de dados
  * Aprender como utilizar a biblioteca SQLC para manusear o banco de dados
  * Fazer o código que realizará a inserção, leitura atualização e exclusão de informações no nosso banco de dados
  * Rodar e testar o código
  
 #### 5º aula - Inserindo valores no banco de dados usando o código da última aula
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
 
 ### 26/01
 
 #### 6º aula - Implementando a transação de banco de dados usando Golang
   * Criar o arquivo store.go que executará as transações
   * Usaremos o princípio ACID - Atomiticy, Consistency, Isolation, Durability. Isso para certificar que:
     * Todas os passos devem ser completos com sucesso. Caso contrário a transação deverá falhar sem alterar o banco de dados.
     * O status do banco de dados deve ser válido após conluir a transação
     * As transações devem ser isoladas e não afetar uma as outras
     * Criar o registro das transações concluídas com sucesso
   * Criar o arquivo store_test.go que usaremos para testar se todas as funções criadas estão funcionando da forma como planejamos
   ##### Passos que serão usados para criar uma transação (Ex: 10 reais da conta 1 para a conta 2):
   * Fazer o registro da transação com valor 10
   * Criar uma movimentação na conta 1 com valor -10
   * Criar uma movimentação na conta 1 com valor +10
   * Retirar 10 da conta 1
   * Adicionar 10 na conta 2

### 27/01

 #### 7º aula - Bloqueio de transação de banco de dados e como lidar com deadlock em Golang
   * Usar o CDD para fazer a atualização do saldo nas contas
   * Observar e aprender como resolver o deadlock 
   * Melhorar as funções que atualizavam a conta para se tornar mais otimizada, juntando-as em apenas uma função que já realiza todo o processo
   
 #### 8º aula - Se aprofundando mais em como evitar deadlock no banco de dados
   * Observando possíveis deadlocks que podem ocorrer no nosso banco de dados que não tinhamos percebido ainda
   * Resolvendo esses deadlocks colocando os processos da atualização da conta na mesma ordem
   * Otimizar esse código para não conter texto duplicado

 #### 9º aula - Entendendo os níveis de isolamento em um banco de dados
   * Entender cada um dos níveis de isolamento, sendo eles:
      * READ UNCOMMITTED
      * READ COMMITTED 
      * REPEATABLE READ 
      * SERIALIZABLE
   * Entender como eles funcionam no mysql
   * Entender como eles funcionam no postgres
     #### Como cada nível de isolamento funciona no MySQL e no Postgres:
     <p float = left>
      <img src='https://github.com/GabrielVitorGL/private/blob/main/Isolation%20Level/MySQL.jpg?raw=true' width='370'>
      <img src='https://github.com/GabrielVitorGL/private/blob/main/Isolation%20Level/Postgres.jpg?raw=true' width='370'>
     </p>
     

 #### 10º aula - Configurar o Github Actions para rodar os unit tests em Go
   * Criar um workflow no Github
   * Escrever o código da nossa action
   * Rodar o código e tratar os erros até que consigamos rodar os testes diretamente do github sem erros

### 28/01

 #### 11º aula - Implementar a API HTTP RESTful no Go usando o Gin
   * Escrever o código dos arquivos conta.go e server.go
   * Escrever o código do arquivo main.go que iniciará o servidor
   * Baixar e utilizar o Postman para visualizar o servidor 
   * Criar os requests para:
|               | Método        | URL do request                                  | Customizações                                                               |
| ------------- | ------------- | ----------------------------------------------- | --------------------------------------------------------------------------- |
| Criar Conta   | POST          | http://localhost:8080/contas                    | Body > raw > JSON > definir o dono da conta e a moeda                       |
| Obter Conta   | GET           | http://localhost:8080/contas/1                  | no lugar de 1 será o ID da conta que você deseja obter                      |
| Listar Contas | GET           | http://localhost:8080/contas?id_pag=1&tam_pag=5 | em 1 você colocará o número da página, e em 5 o número de contas por página |
