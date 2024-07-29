# Go Ethereum Fetcher

An application to interact with the Ethereum blockchain and manage transactions using Go and the Gin framework.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Endpoints](#endpoints)
- [Code Overview](#code-overview)
  - [Handlers](#handlers)
  - [Repositories](#repositories)
  - [Ethereum Gateway](#ethereum-gateway)
  - [Models](#models)
- [Running Tests](#running-tests)
- [License](#license)

## Prerequisites

- Go 1.18 or higher
- Docker
- Docker Compose

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/gsstoykov/go-ethereum-fetcher.git
   cd go-ethereum-fetcher
   ```

2. Build and start the application using Docker:
   ```sh
   docker-compose build
   docker-compose up
   ```

## Configuration

The application requires a few environment variables to be set. Create a `.env` file in the root of your project with the following content:

```sh
API_PORT=<your_api_port>
DB_CONNECTION_STRING=<your_db_con_str>
ETH_NODE_URL=<url_with_your_api_key>
JWT_STRING=<your_secret>
PRIVATE_KEY=<your_private_key>
CONTRACT_ADDRESS=<your_contract_address>
WS_NODE_URL=<ws_url_with_your_api_key>
```

## Endpoints

### Fetch Users

- **Endpoint:** `GET /users`
- **Description:** Retrieves a list of all users.
- **Request:**
  - **Method:** GET
  - **URL:** `/users`
- **Response:**
  - **Body:**
    ```json
    {
      "users": [
        {
          "id": 1,
          "username": "johndoe",
          "created_at": "2024-07-24T10:00:00Z",
          "updated_at": "2024-07-24T10:00:00Z"
        }
      ]
    }
    ```
  - **Description:** Returns a list of users with their details.

### Create User

- **Endpoint:** `POST /user`
- **Description:** Creates a new user.
- **Request:**
  - **Method:** POST
  - **URL:** `/user`
  - **Body:**
    ```json
    {
      "username": "johndoe",
      "password": "password123"
    }
    ```
- **Response:**
  - **Body:**
    ```json
    {
      "user": {
        "id": 1,
        "username": "johndoe",
        "created_at": "2024-07-24T10:00:00Z",
        "updated_at": "2024-07-24T10:00:00Z"
      }
    }
    ```
  - **Description:** Returns the created user’s details.

### Authenticate

- **Endpoint:** `POST /auth`
- **Description:** Authenticates a user.
- **Request:**
  - **Method:** POST
  - **URL:** `/auth`
  - **Body:**
    ```json
    {
      "username": "johndoe",
      "password": "password123"
    }
    ```
- **Response:**
  - **Body:**
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNjI4Mjg3MDk3LCJleHBpcmF0aW9uIjoiZXhhbXBsZXMiLCJsaWIiOiJleGFtcGxlIn0.4rOwh0zeF8D90L8JmLHKSH3xnlUu7rd9R6p5HDch4Q8"
    }
    ```
  - **Description:** Returns a JWT token if authentication is successful.

### Fetch User Transactions

- **Endpoint:** `GET /my`
- **Description:** Retrieves transactions for the authenticated user.
- **Request:**
  - **Method:** GET
  - **URL:** `/my`
  - **Headers:**
    - `Authorization: Bearer <token>`
- **Response:**
  - **Body:**
    ```json
    {
      "transactions": [
        {
          "id": 1,
          "hash": "0x123...",
          "amount": "1000",
          "timestamp": "2024-07-24T10:00:00Z"
        }
      ]
    }
    ```
  - **Description:** Returns a list of queried transactions for the authenticated user.

## Transaction Routes

### Fetch Transactions

- **Endpoint:** `GET /transactions`
- **Description:** Retrieves a list of all transactions.
- **Request:**
  - **Method:** GET
  - **URL:** `/transactions`
- **Response:**
  - **Body:**
    ```json
    {
      "transactions": [
        {
          "id": 1,
          "hash": "0x123...",
          "amount": "1000",
          "timestamp": "2024-07-24T10:00:00Z"
        }
      ]
    }
    ```
  - **Description:** Returns a list of transactions with their details.

### Fetch Transactions List

- **Endpoint:** `GET /eth`
- **Description:** Retrieves transaction details based on the provided transaction hashes. Optional authentication.
- **Request:**
  - **Method:** GET
  - **URL:** `/eth?transactionHashes=<hash1>&transactionHashes=<hash2>`
  - **Headers:**
    - `Authorization: Bearer <token>`
- **Response:**
  - **Body:**
    ```json
    {
      "transactions": [
        {
          "id": 1,
          "hash": "0x123...",
          "amount": "1000",
          "timestamp": "2024-07-24T10:00:00Z"
        }
      ]
    }
    ```
  - **Description:** Returns details for the specified transaction hashes.

## Person Routes

### Save Person

- **Endpoint:** `POST /savePerson`
- **Description:** Attempts to save a new PersonInfo on deployed SimplePersonInfoContract on the Ethereum network.
- **Request:**
  - **Method:** POST
  - **URL:** `/savePerson`
  - **Headers:**
    - `Authorization: Bearer <token>`
  - **Body:**
    ```json
    {
      "name": "Jane Doe",
      "age": 30
    }
    ```
- **Response:**
  - **Body:**
    ```json
    {
      "tx": {
        "hash": "0x123...",
        "status": "pending"
      }
    }
    ```
  - **Description:** Returns the call transaction and its status(usually pending as state does not change instantly).

### List People

- **Endpoint:** `GET /listPeople`
- **Description:** Retrieves a list of all people handled by the smart contract and saved to app db.
- **Request:**
  - **Method:** GET
  - **URL:** `/listPeople`
- **Response:**
  - **Body:**
    ```json
    {
      "people": [
        {
          "id": 1,
          "name": "Jane Doe",
          "age": 30
        }
      ]
    }
    ```
  - **Description:** Returns the list of people with their details saved to db.

## Running Tests

Tests can be run using the Go testing framework. Use the following command to run all tests:

```sh
go test ./...
```

## License

This project is licensed under the MIT License.
