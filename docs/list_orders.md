# List Orders Feature Documentation

## Overview
The "List Orders" feature allows clients to retrieve a list of all existing orders in the system. This feature is implemented following Clean Architecture principles and is exposed via three different interfaces: REST, gRPC, and GraphQL.

## Architecture
The feature traverses the following layers:
1.  **Entity Layer**: Defines the `OrderRepositoryInterface` with the `List()` method.
2.  **Use Case Layer**: `ListOrdersUseCase` orchestrates the retrieval of data from the repository and formats it into `ListOrdersOutputDTO`.
3.  **Infrastructure Layer**:
    *   **Database**: `OrderRepository` implements the `List()` method using SQL queries.
    *   **Web (REST)**: `WebOrderHandler` handles HTTP GET requests.
    *   **gRPC**: `OrderService` handles gRPC `ListOrders` requests.
    *   **GraphQL**: `Resolver` handles the `orders` query.

## API Specifications

### 1. REST API
*   **Endpoint**: `GET /order/list`
*   **Description**: Retrieves all orders.
*   **Response Format**: JSON
*   **Success Response (200 OK)**:
    ```json
    [
      {
        "id": "string",
        "price": 100.0,
        "tax": 10.0,
        "final_price": 110.0
      },
      ...
    ]
    ```
*   **Error Response (500 Internal Server Error)**:
    ```json
    "error message"
    ```

### 2. gRPC API
*   **Service**: `OrderService`
*   **Method**: `ListOrders`
*   **Request**: `Blank` (Empty message)
*   **Response**: `ListOrdersResponse`
    ```protobuf
    message ListOrdersResponse {
      repeated CreateOrderResponse orders = 1;
    }
    
    message CreateOrderResponse {
      string id = 1;
      float price = 2;
      float tax = 3;
      float final_price = 4;
    }
    ```

### 3. GraphQL API
*   **Query**: `orders`
*   **Description**: Fetches a list of orders.
*   **Schema Definition**:
    ```graphql
    type Query {
      orders: [Order!]!
    }
    
    type Order {
      id: String!
      Price: Float!
      Tax: Float!
      FinalPrice: Float!
    }
    ```
*   **Example Query**:
    ```graphql
    query {
      orders {
        id
        Price
        FinalPrice
      }
    }
    ```

## Implementation Details
*   **Dependency Injection**: Dependencies are managed in `cmd/ordersystem/main.go` (and `wire.go` for future generation).
*   **Database**: Uses SQLite (or configured DB) to query the `orders` table.
