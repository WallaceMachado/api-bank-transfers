






> status:	🚧 api-bank-transfers 🚀 em construção..  🚧

API, que responde JSON, de transferencia entre contas Internas de um banco digital.


## Indice

* <p><a href="#pré-requisitos">Pré Requisitos</a> </p>
* <p><a href="#iniciando-projeto">Iniciando Projeto</a></p>
* <p><a href="#variáveis-de-ambiente">Variáveis de Ambiente</a></p>
* <p><a href="#rotas">Rotas</a></p>
* <p><a href="#controle-de-versão">Controle de versão</a></p>
* <p><a href="#testes">Testes</a></p>
* <p><a href="#autor">Autor</a></p>




## Pré Requisitos

Antes de começar, você precisará ter as seguintes ferramentas instaladas em sua máquina:
* [Git](https://git-scm.com)

Para rodar via docker
* [Docker](https://docs.docker.com/)

Para rodar Local
* [Go](https://golang.org/) versão 1.16.7
* [Postgres](https://www.postgresql.org/)

Além disso, é bom ter um editor para trabalhar com o código como: [VSCode](https://code.visualstudio.com/)



## Iniciando Projeto 

### Local

Deverá criar o banco de dados com as configurações informadas no arquivo ``` .env ``` DB_USER, DB_NAME, DB_PASS

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/api-bank-transfers.git

# Acesse a pasta do projeto no terminal / cmd
$ cd api-bank-transfers

# Instale as dependências e rode o projeto
$ go run main.go

# Server is running
```
### Docker

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/api-bank-transfers.git

# Acesse a pasta do projeto no terminal / cmd
$ cd api-bank-transfers

# Instale as dependências e rode o projeto
$ docker-compose up --build

```

## Variáveis de Ambiente

Após clonar o repositório, renomeie o ``` .env.example ``` no diretório raiz para ``` .env ``` e atualize com suas configurações.


| Chave  |  Descrição  | Predefinição  |
| :---: | :---: | :---: | 
|  PORT |  Número da porta em que o aplicativo será executado. | 5000  |
|  DB_HOST |  Host Postgres.  | pg  |
|  DB_PORT |  Porta Postgres.  |  5432  |
|  DB_USER |  Usuário Postgres. |  -  |
|  DB_NAME |  Nome do banco de dados do aplicativo. |  -  |
|  DB_PASS |  Senha do Postgres.  |  -   |
|  DB_TYPE | tipo do banco de dados.  |  postgres  |
|  SECRET_KEY_JWT | Uma string alfanumérica aleatória. Usado para criar tokens assinados.  |  -   |

## Controle de versão
Para contrele de versão, foi inserida a versão ``` v1 ``` após o  ``` host ```

```
GET http://localhost:5000/api/v1/transfers

```


## Rotas

| Rotas  |  HTTP Method  | Params  |  Descrição  |  Auth Method  |
| :---: | :---: | :---: | :---: | :---: |
|  /accounts |  POST |  Body: ``` name ```, ``` cpf ```, ``` secret ``` e ``` balance ``` |  Crie uma nova conta |  ❌ |
|  /accounts |  GET |  -  | Recupere uma lista com todas as contas |  ❌ |
|  /accounts/:account_id/balance |  GET |  Params: ``` account_id ``` |  Consulte o saldo de uma conta |  ❌ |
|  /login |  POST | -  |  Faça login  |  ❌ |
|  /transfers |  POST |  Body: ``` account_destination_id ``` e ``` amount ```   |  Faça uma transferência bancária |  Bearer |
|  /transfers |  GET |  -  |  Consulte as transferências de uma conta |  Bearer |

Rotas com Bearer como método de autenticação esperam um cabeçalho de autorização. Consulte a seção [Bearer Token](#bearer-token) para mais informações.

### Requisições
* ``` POST /accounts ```

Corpo da requisição:
  
```
{
    "name": "teste",
    "cpf": "31410884031",
    "secret": "123456",
    "balance": 1000
}

```

* ``` POST /login ```

Corpo da requisição:
  
```
{
    "cpf": "31410884031",
    "secret": "123456"
}

```

* ``` POST /transfers ```

Corpo da requisição:
  
```
{
    "account_destination_id":"4b808c0b-5822-410d-a67d-d87241f03e9d",
    "amount":100
}

```

  
## Bearer Token
Algumas rotas esperam um Bearer Token em um cabeçalho de autorização.


> Você pode ver essas rotas na seção de [Rotas](#rotas).

```
GET http://localhost:5000/api/v1/transfers Authorization: Bearer <token>
```
>Para obter este token, você só precisa se autenticar por meio da rota ``` /login ``` e ela retornará a chave do token com um Bearer Token válido


## Testes
Para executar os testes :

### Local

Altere a varável de ambente DB_HOST para ``` localhost ```

```bash

  # Acesse a pasta do projeto no terminal / cmd
  $ cd api-bank-transfers
  
  # Rode os testes
  $ go test ./...
  
```

### Docker


```bash

    # Acesse a pasta do projeto no terminal / cmd
    $ cd api-bank-transfers

    # Instale as dependências e rode o projeto
    $ docker-compose up --build

    # Pegar o CONTNAINER ID  da imagem github.com/wallacemachado/api-bank-transfers
    $ docker ps
    # Exemplo de retorno:
    # CONTAINER ID      IMAGE                                             COMMAND                  
    # xxxxxxxxxxxxxxx   github.com/wallacemachado/api-bank-transfers      "go run main.go"         0.0.0.0:5003->5000/tcp, :::5003->5003 ....  
    # yyyyyyyyyyyyyyy   github.com/wallacemachado/api-bank-transfers-db   "docker-entrypoint.s…"   0.0.0.0:5432->5432/tcp, :::5432->5432 ....

    # acesse o contaniner
    $ docker exec -it <CONTAINER ID> /bin/bash
    # Exemplo:
    # docker exec -it xxxxxxxxxxxxxxx /bin/bash

    #root@xxxxxxxxxxxxxxx:/go/src/app# 
    #rode os testes
    $ go test ./...

```



## Autor


Feito com ❤️ por [Wallace Machado](https://github.com/WallaceMachado) 🚀🏽 Entre em contato!

[<img src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/in/wallace-machado-b2054246/)
