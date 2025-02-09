# go-backend-demo

## Overview
`go-backend-demo` is a demonstration backend API that simulates real-world API use cases. This backend is designed to generate Swagger UI documentation and will be used by the [swagger-mcp](https://github.com/danishjsheikh/swagger-mcp) project.

This backend takes reference from the following repositories:
- [MehmetFiratKomurcu - golang-swaggo](https://github.com/MehmetFiratKomurcu/Youtube-Channel/tree/master/golang-swaggo)
- [joefazee - Go API Documentation with Swagger](https://github.com/joefazee/go-api-doc-with-swagger/tree/main)

## Features
- RESTful API with endpoints for users and orders
- Swagger UI documentation generation
- Demonstrates authentication and request handling

## Installation and Setup
To run this project locally, follow these steps:

1. Install dependencies:
   ```sh
   go mod download
   go mod tidy
   ```

2. Generate Swagger documentation:
   ```sh
   swag init
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints
### Base Configuration
- **Host:** `localhost:8080`
- **Base Path:** `/v1`

### Orders API
#### Create Order
- **Endpoint:** `POST /orders`
- **Description:** Creates an order with the provided request.
- **Headers:**
  - `x-correlationid` (string, required) - Code of the order.
- **Request Body:**
  ```json
  {
    "age": 25,
    "countryCode": "US",
    "shipmentNumber": "123456789"
  }
  ```
- **Responses:**
  - `200 OK`: Order successfully created.
  - `400 Bad Request`: Invalid request.

#### Get Order by Code
- **Endpoint:** `GET /orders/{orderCode}`
- **Description:** Retrieves order details by code.
- **Headers:**
  - `x-correlationid` (string, required) - Code of the order.
- **Path Parameters:**
  - `orderCode` (string, required) - The order code.
- **Responses:**
  - `200 OK`: Order details returned.

### Users API
#### Get All Users
- **Endpoint:** `GET /users`
- **Description:** Retrieves a list of all users.
- **Responses:**
  - `200 OK`: List of users returned.

#### Create User
- **Endpoint:** `POST /users`
- **Description:** Creates a new user.
- **Security:** Requires a Bearer token.
- **Request Body:**
  ```json
  {
    "username": "johndoe",
    "email": "john@example.com",
    "full_name": "John Doe",
    "is_active": true
  }
  ```
- **Responses:**
  - `201 Created`: User successfully created.
  - `400 Bad Request`: Invalid request.

## Security
- Uses API key-based authentication with a Bearer token.
- Security Definition:
  ```json
  {
    "type": "apiKey",
    "name": "Authorization",
    "in": "header"
  }
  ```

## Swagger Documentation
Swagger UI will be available at:
```
http://localhost:8080/swagger/index.html
```
