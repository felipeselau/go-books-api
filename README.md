# go-books-api

API REST de livros desenvolvida em Go, utilizando o framework Gin e GORM para operações de CRUD de livros.

## Descrição

Este projeto implementa uma API REST simples para gerenciamento de livros, permitindo operações de criação, leitura, atualização e exclusão (CRUD). A aplicação está estruturada com rotas específicas para cada operação de livro e usa GORM para interação com o banco de dados PostgreSQL.

## Estrutura do Projeto

- **main.go**: Configura a aplicação, carregando variáveis de ambiente e iniciando a conexão com o banco de dados.
- **server/server.go**: Define e inicia o servidor, configurando as rotas principais da API.
- **routes/router.go**: Configura as rotas para o endpoint `/api/v1/books`.
- **models/book.go**: Define o modelo `Book` usado no banco de dados.
- **database/database.go**: Gerencia a conexão e configuração do banco de dados.
- **database/migrations/migrations.go**: Executa as migrações para o banco de dados.
- **controllers/book_controller.go**: Implementa as funções de CRUD para os endpoints da API.

## Instalação e Uso

1. Clone o repositório:
   ```bash
   git clone https://github.com/felipeselau/go-books-api.git
   ```

2. Navegue até a pasta do projeto:
   ```bash
   cd go-books-api
   ```

3. Configure suas variáveis de ambiente criando um arquivo `.env`:
   ```
   DATABASE_URL=postgres://seu_usuario:senha@localhost:5432/seu_banco
   ```

4. Instale as dependências:
   ```bash
   go mod tidy
   ```

5. Execute o projeto:
   ```bash
   go run main.go
   ```

A aplicação estará disponível em `http://localhost:5000/api/v1/books`.

## Endpoints

- `GET /api/v1/books/`: Lista todos os livros.
- `POST /api/v1/books/`: Cria um novo livro.
- `GET /api/v1/books/:id`: Retorna os detalhes de um livro específico.
- `PUT /api/v1/books/:id`: Atualiza as informações de um livro.
- `DELETE /api/v1/books/:id`: Deleta um livro.

## Exemplo de JSON

Para criar ou atualizar um livro, envie um JSON com a seguinte estrutura:

```json
{
  "name": "Example Book",
  "description": "An example description.",
  "medium_price": 19.99,
  "author": "John Doe",
  "image_url": "http://example.com/image.jpg"
}
```

## Tecnologias

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
