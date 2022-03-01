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
      <img src='https://github.com/GabrielVitorGL/simplebank/blob/main/images/Isolation%20Level/MySQL.jpg?raw=true' width='370'>
      <img src='https://github.com/GabrielVitorGL/simplebank/blob/main/images/Isolation%20Level/Postgres.jpg?raw=true' width='370'>
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
   * Criar os seguintes requests:
 
   | Função        | Método        | URL do request                                  | Customizações                                                               |
   | ------------- | ------------- | ----------------------------------------------- | --------------------------------------------------------------------------- |
   | Criar Conta   | POST          | http://localhost:8080/contas                    | Body > raw > JSON > definir o dono da conta e a moeda                       |
   | Obter Conta   | GET           | http://localhost:8080/contas/1                  | no lugar de 1 será o ID da conta que você deseja obter                      |
   | Listar Contas | GET           | http://localhost:8080/contas?id_pag=1&tam_pag=5 | em 1 você colocará o número da página, e em 5 o número de contas por página |
   
### 29/01

 #### 12º aula - Criar e carregar um arquivo config & Variáveis de ambiente em Golang usando o Viper
   * Fazer um arquivo.env que conterá as configurações desejadas
   * Criar o arquivo config.go
   * Usar o Viper para carregar essas configurações
   * Iniciar o servidor e fazer os testes no Postman para checar se tudo funcionou

### 30/01

 #### 13º aula - Simulando o Banco de dados para testar a API HTTP e obter 100% de cobertura no teste
   * Simular o banco de dados nos trará diversas vantagens, podemos citar:
       * Rodar testes independentes um dos outros, ou seja, um não irá interferir no outro e não ficará gravado no banco de dados
       * Rodas os testes mais rápido, por não presicar se conectar com o banco de dados e tudo mais
       * 100% de cobertura no teste, e também podemos testar diversos tipos de erros inesperados como perda de conexão etc
   * Utilizar a biblioteca Gomock para simular o nosso banco de dados 
   * Fazer o código que irá testar 100% das nossas funções

### 05/02

 #### 14º aula - Implementando a API de transferência de dinheiro com um validador de parâmetros customizado
   * Criar o arquivo transferencia.go na pasta api que será responsável por fazer a transferencia de dinheiro entre contas
   * Testar no Postman se a API está funcionando
  
   | Função              | Método | URL do request                       | Customizações                                                           |
   | ------------------- | ------ | ------------------------------------ | ----------------------------------------------------------------------- |
   | Criar Transferencia | POST   | http://localhost:8080/transferencias | Body > raw > JSON > definir: de_id_conta, para_id_conta, quantia, moeda |
   
   * Retirar a opção "oneof" de nossas API e fazer essa verificação de uma forma mais otimizada, e com menos chances de cometer erros caso fossemos trabalhar com centenas de tipos de moedas

### 10/02

 #### 15º aula - Criando a tabela de usuários com uma chave única e estrangeira no PostgreSQL
   * Implementaremos o recurso de autenticação e autorização de usuário
   * Criar a tabela usuários no dbdiagram.io
   * Atualizar o nosso banco de dados com a nova tabela
       * Obs: Poderíamos substituir todo o código do arquivo migrateup com o novo, deletar o banco de dados e iniciar outro com as mudanças, porém essa não é a maneira correta visto que em um projeto real temos diversas atualizações que podem ser realizadas e utilizar esse processo iria apagar todos os dados salvos em nosso banco de dados. Desta forma, iremos criar outro arquivo migrate para fazer essa atualização
   * Testar os códigos que criamos para ver se está tudo funcionando
 
 #### 16º aula - Atualizando o código em Go para funcionar com a tabela criada na última aula
   * Criar e fazer o código do arquivo usuario.sql que irá funcionar com a tabela "usuario" que criamos na aula 15
   * Fazer os testes com essa nova implementação
   * Tratar os erros gerados com essa nova tabela
   * Testar e arrumar os requests no Postman

### 11/02

 #### 17º aula - Armazenando senhas de forma segura usando o Bcrypt
   * Nunca devemos armazenar senhas em nosso banco de dados, primeiro devemos encriptar esse dado para, após isso, podermos armazená-lo com segurança

<img src='https://github.com/GabrielVitorGL/simplebank/blob/main/images/Hash%20Password/Hash.jpg?raw=true' style="width: 50%;">

   * Fazer as funções que encriptam e decriptam a senha
   * Fazer os testes para garantir que as funções estão funcionando corretamente
   * Implementar essas funções no teste de criação de usuário
   * Codificar a API de criação de usuários
   * Testar a API
   
   | Função        | Método | URL do request                 | Customizações                                                          |
   | --------------| ------ | ------------------------------ | ---------------------------------------------------------------------- |
   | Criar Usuario | POST   | http://localhost:8080/usuarios | Body > raw > JSON > definir: nome_usuario, senha, nome_completo, email |
   
   * Retirar a senha encriptada do response da API, já que o usuário não precisa dessa informação e mostrá-la pode ser um problema

 #### 18º aula - Escrevendo testes de unidade mais fortes com um matcher gomock personalizado
   * Criar o arquivo usuario_test para testar a API
   * Verificamos que como a senha sempre será diferente cada vez que encriptamos-a, não podemos usar uma simples função de checar valores correspondentes. Então, iremos criar uma função específica para isso
   * Realizar os testes e certificar que tudo está funcionando

 #### 19º aula - Autenticação baseada em token: os problemas de segurança do JWT e utilizando o PASETO para resolver esses problemas
   * Algumas desvantagens de usar o JWT:
      * Algoritmos fracos: Temos a opção de escolher diversos algorítimos para trabalhar, entre eles vários não tão seguros. Dar essa liberdade de escolha ao usuário pode ser um problema
      * Falsificação trivial: Algumas escolhas de bibliotecas e outros descuidos podem abrir brechas de segurança em nosso sistema, como um exemplo já ocorrido de uma falha em que era possível alterar o "alg" do header para "none" ou "HS256"
   * Como uma melhor e mais segura alternativa de autenticação, temos o PASETO
      * Algoritmos fortes: Os desenvolvedores não tem que escolher qual algorítimo o sistema irá utilizar, tornando assim muito mais simples e garantido que estaremos sempre utilizando o nível de segurança máxima. A única coisa que temos que escolher é a versão do PASETO que iremos utilizar
      * Anti Falsificação trivial: Nesse caso como não temos mais o "alg" no header, isso previne os ataques do tipo "none"
      * O payload não é codificado como no JWT, e sim encriptado. Isso torna muito mais seguro pois dessa forma não é possivel ler ou mudar os dados armazenados nele

### 12/02

 #### 20º aula - Criar e verificar os tokens JWT e PASETO usando Golang
   * Criar o arquivo maker.go e payload.go que definirão as estruturas do token
   * Implementar o JWT e testar para ver como se comporta
   * Implementar o PASETO e testar
   * Observar a diferença entre os dois métodos de verificação de token

### 14/02

 #### 21º aula - Implementando a API de Login de usuário que retorna um token PASETO ou JWT
   * Definir as funções que criarão o token na API
   * Implementar a API de login e corrigir os erros
   * Testar a API
   
   | Função        | Método | URL do request                       | Customizações                                    |
   | --------------| ------ | ------------------------------------ | ------------------------------------------------ |
   | Logar Usuario | POST   | http://localhost:8080/usuarios/login | Body > raw > JSON > definir: nome_usuario, senha |
   
### 16/02

 #### 22º aula - Implementando regras de autenticação middleware
   * Criar o arquivo middleware.go
   * Fazer os unit tests para verificar se tudo está funcionando como queremos
   * Implementar o middleware no servidor
   * Implementar as regras de autenticação, que definirão as permissões dos usuários relacionadas a cada função

<img src='https://github.com/GabrielVitorGL/simplebank/blob/main/images/Authorization%20Rules/Rules.jpg?raw=true' style="width: 50%;">

   * Testar as regras no postman

### 26/02

 #### 23º aula - Criando uma imagem do Docker pequena da nossa aplicação em Golang
   * Criar uma branch no github para fazer o processo de criação da imagem
   * Criar a imagem de nosso projeto com o Dockerfile
   * Reduzir o tamanho da imagem, compilando apenas o necessário. 
      * A redução será de mais de 500mb para apenas 22mb


 #### 24º aula - Conectando dois containers na mesma rede do docker
   * Rodar o server usando a imagem que criamos e arrumar os erros
   * Criar uma network para colocar os containers do postgres e do simplebank, para que possam se achar na mesma rede apenas pelo nome

 #### 25º aula - Fazendo o arquivo docker-compose e controlando os pedidos de inicialização usando wait-for.sh
   * Fazer o arquivo docker-compose.yaml que irá conter todas as configurações dos containers
   * Fazer o app esperar o servidor do postgres estar pronto antes de rodar o migrate script
   * Testar para conferir se o servidor está funcionando corretamente

 #### 26º aula - Criando uma conta gratuita no Amazon Web Services
   * Criar uma conta gratuita no serviço de processamento em nuvem da amazon para usarmos com o nosso projeto

### 28/02

 #### 27º aula - Criando e enviando a imagem docker automaticamente para o AWS usando o github actions
   * Usar o Amazon ECR para criar um repositório no AWS para nosso projeto
   * Fazer o arquivo deploy.yml que conterá os passos para automatizar o processo no github actions
   * Utilizar o IAM para obtermos as acess keys
   * Criar um grupo com as permissões necessárias para o github conseguir acessar os serviços da amazon
   * Criptografar as duas acess keys que teremos que colocar no arquivo, pois não podemos colocar diretamente no arquivo por questões de segurança. Utilizaremos a funcionalidade "secrets" do próprio github
   * Testar se tudo está funcionando e a action está criando e enviando a imagem automaticamente para o ECR
 

 #### 28º aula - Criando um Banco de dados no AWS usando o RDS
   * Criar o banco de dados do nosso projeto no AWS
   * Tornar possível conectar a esse banco de dados em qualquer lugar
   * Conectar e executar o comando para criar as tabelas nesse banco de dados

 #### 29º aula - Usando o AWS Secrets Manager para armazenar e utilizar segredos criptografados
   * Criar as chaves para nossas variáveis no Secrets Manager
   * Instalar e configurar o AWS CLI no nosso computador
   * Executar o comando que criará o arquivo com as chaves que definimos antes
   * Utilizar o JQ para organizarr o arquivo com as chaves que o AWS irá criar
   * Fazer com que o github actions substitua essas variáveis para nós
   * Baixar a imagem para ver se há erros ou não
