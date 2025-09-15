# BookApiGo

API REST para gerenciamento de livros desenvolvida em Go, utilizando o framework Gin e GORM como ORM.

## Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- Git

## Configuração do Projeto

### Clonando o repositório

```sh
git clone https://github.com/Kovokar/BookApiGo.git
cd BookApiGo
```

### Instalando dependências

```sh
go mod download
go mod tidy
```

### Configurando o ambiente

1. Inicie o banco de dados PostgreSQL:
```sh
docker-compose up -d
```

2. Execute a aplicação:
```sh
go run cmd/run.go
```

A API estará disponível em `http://localhost:8081/api/v1`

## Estrutura do Projeto

```
bookApiGo/
├── cmd/
│   └── run.go
├── internal/
│   ├── handlers/
│   ├── models/
│   └── routes/
├── docker-compose.yml
└── README.md
```

## Endpoints da API

### Livros

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/api/v1/books` | Lista todos os livros |
| GET | `/api/v1/books/:id` | Busca um livro por ID |
| POST | `/api/v1/books` | Cria um novo livro |
| POST | `/api/v1/books/bulk` | Cria múltiplos livros |
| PUT | `/api/v1/books/:id` | Atualiza um livro |
| DELETE | `/api/v1/books/:id` | Remove um livro |

### Exemplos de Requisições

#### Criar um livro
```json
POST /api/v1/books
{
    "name": "O Senhor dos Anéis",
    "description": "Uma história épica de fantasia",
    "medium_price": 89.90,
    "author": "J.R.R. Tolkien",
    "image_url": "https://exemplo.com/imagem.jpg"
}
```

#### Criar múltiplos livros
```json
POST /api/v1/books/bulk
[
    {
        "name": "Livro 1",
        "description": "Descrição 1",
        "medium_price": 29.90,
        "author": "Autor 1",
        "image_url": "https://exemplo.com/imagem1.jpg"
    },
    {
        "name": "Livro 2",
        "description": "Descrição 2",
        "medium_price": 39.90,
        "author": "Autor 2",
        "image_url": "https://exemplo.com/imagem2.jpg"
    }
]
```

## Desenvolvimento

### Executando testes

```sh
go test ./...
```

## Observações

- O banco de dados PostgreSQL é configurado através do `docker-compose.yml`
- As migrações são executadas automaticamente na inicialização
- A API utiliza o padrão RESTful
- Todas as respostas são em formato JSON

## Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais