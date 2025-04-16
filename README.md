# TaskFlow API

A **TaskFlow API** é uma aplicação desenvolvida com **Gin** (Go) para gerenciar tarefas (tasks). Ela usa **PostgreSQL** como banco de dados e **UUID** como identificador único das tarefas. O projeto é configurado para ser executado em **Docker**.

## Tecnologias

- **Gin** - Framework web para Go
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **Docker** - Para conteinerização da aplicação e banco de dados
- **UUID** - Identificador único das tarefas

## Pré-requisitos

Antes de rodar o projeto, você precisará ter as seguintes ferramentas instaladas:

- **Docker**: [Instalar o Docker](https://www.docker.com/get-started)
- **Docker Compose**: [Instalar o Docker Compose](https://docs.docker.com/compose/install/)

## Rodando o Projeto

### Passo 1: Clonar o Repositório

Se você ainda não fez isso, clone o repositório para a sua máquina local:

```bash
git clone https://github.com/seu-usuario/taskflow.git
cd taskflow
```

### Passo 2: Configuração do Banco de Dados

O PostgreSQL já está configurado no docker-compose.yml para rodar com a imagem postgres:13. O banco de dados será chamado taskdb.

Um arquivo init.sql é montado no contêiner para garantir que a extensão uuid-ossp seja ativada no PostgreSQL. Esta extensão é necessária para gerar UUIDs automaticamente para os IDs das tarefas.

### Passo 3: Configurar Variáveis de Ambiente
As variáveis de ambiente necessárias para a conexão com o banco de dados já estão configuradas no docker-compose.yml
```yaml
    version: "3.8"
    services:
    db:
        image: postgres:13
        container_name: postgres_db
        environment:
        POSTGRES_USER: user
        POSTGRES_PASSWORD: password
        POSTGRES_DB: taskdb
        ports:
        - "5432:5432"
        volumes:
        - postgres_data:/var/lib/postgresql/data
        - ./init.sql:/docker-entrypoint-initdb.d/init.sql  # Montando o script de inicialização
        networks:
        - task_network

    app:
        build:
        context: .
        container_name: task_app
        environment:
        - DB_HOST=db
        - DB_USER=user
        - DB_PASSWORD=password
        - DB_NAME=taskdb
        ports:
        - "8080:8080"
        depends_on:
        - db
        networks:
        - task_network

    volumes:
    postgres_data:
        driver: local

    networks:
    task_network:
        driver: bridge
```

### Passo 4: Iniciar o Docker Compose
```bash
docker-compose up --build -d
```
Isso fará o seguinte:

Criar e iniciar o contêiner do banco de dados PostgreSQL.

Criar e iniciar o contêiner da aplicação Go.

Rodar as migrações para criar a tabela de tarefas com UUIDs.

### Observações: 

Verificar Logs do Docker:
```bash
docker-compose logs -f
```
Parar os Contêineres:
```bash
docker-compose down
```
Se você quiser remover os volumes de dados também (o que irá apagar os dados do banco), execute:
```bash
docker-compose down -v
```
