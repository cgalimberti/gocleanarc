# Go Expert - Clean Architecture Order System

Este projeto é um exemplo de implementação de Clean Architecture em Go, incluindo APIs REST, GraphQL e gRPC.

## Pré-requisitos

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/) (1.20+)

## Como rodar a aplicação

### 1. Subir o ambiente Docker

O projeto utiliza Docker Compose para subir o banco de dados MySQL e o RabbitMQ.

```bash
docker-compose up -d
```

### 2. Configurar o Banco de Dados

Acesse o container do MySQL e crie a tabela necessária.

```bash
docker exec -it mysql mysql -uroot -proot orders
```

Execute o seguinte comando SQL para criar a tabela `orders`:

```sql
CREATE TABLE orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
);
```

### 3. Rodar a aplicação

Na raiz do projeto, execute o comando:

```bash
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```

A aplicação iniciará os seguintes servidores:
- **Web Server (REST)**: Porta 8000
- **gRPC Server**: Porta 50051
- **GraphQL Server**: Porta 8080

## Como testar

### REST API

**Criar Order**
```bash
curl -X POST http://localhost:8000/order -H "Content-Type: application/json" -d '{
    "id": "uuid-test-1",
    "price": 100.5,
    "tax": 0.5
}'
```

**Listar Orders**
```bash
curl http://localhost:8000/order/list
```

### GraphQL

Acesse o Playground em: [http://localhost:8081/](http://localhost:8081/)

**Query para listar orders:**
```graphql
query {
  orders {
    id
    Price
    FinalPrice
  }
}
```

**Mutation para criar order:**
```graphql
mutation {
  createOrder(input: {id: "uuid-test-graphql", Price: 200.0, Tax: 2.0}) {
    id
    Price
    FinalPrice
  }
}
```

### gRPC

Para testar o gRPC, você pode utilizar uma ferramenta como o [Evans](https://github.com/ktr0731/evans).

```bash
evans -r repl
```

Dentro do REPL do Evans:
1.  Selecione o package: `package pb`
2.  Selecione o service: `service OrderService`
3.  Chame o método: `call CreateOrder` ou `call ListOrders`
