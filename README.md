# Go Ethereum Fetcher

An application to interact with the Ethereum blockchain and manage transactions using Go and the Gin framework.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Architecture](#architecture)
- [Endpoints](#endpoints)
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
   docker compose build
   docker compose up
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

### [Opt.] Deploy `SimplePersonInfoContract.sol`

Smart contract ABI was generated following:
https://goethereumbook.org/smart-contracts/
https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings

```sh
go run contract/deploy/person.go
```

For testing purposes you can use my already deployed contract found at:
https://sepolia.etherscan.io/address/0xEe0D53C64AC1aad09861a139c52bDD087d3eeaCC
On this address you can also see my call history during development.

## Architecture

![Alt text](xdiagrams/architecture_overview.png)

### Fetcher Overview

The Ethereum `Fetcher` is a REST API that facilitates interaction with Ethereum blockchain data. It provides endpoints for user management, transaction retrieval, and data persistence via database storage, utilizing middleware for authentication and data processing.

![Alt text](xdiagrams/fetcher_no_contract.png)

#### Handle Manager

The `HandleManager` responsible for setting up and managing the HTTP routes and their respective handlers using the Gin web framework. This component interacts with various other components such as repositories, handlers, middleware, and the Ethereum client to serve API requests.

#### Handles

##### Transaction Handle

The `TransactionHandler` is responsible for handling HTTP requests related to transactions. It works closely with the transaction and user repositories, as well as the `Ethereum Gateway`, to process and respond to transaction-related operations.

##### User Handle

The `UserHandler` is responsible for handling HTTP requests related to user operations such as user creation, authentication, fetching user details, and fetching user transactions. It integrates with the user repository and employs JWT for secure authentication.

#### Middleware

##### Authentication Middleware

The `AuthenticateMiddleware` is responsible for handling JWT token authentication for HTTP requests. It ensures that only authenticated users can access certain routes by validating the JWT token provided in the Authorization header of the request.

##### Decode Middleware

The `RLPDecodeMiddleware` is responsible for decoding specific RLP transaction hash input in HTTP requests.

#### Ethereum Gateway

The `EthereumGateway` provides a mechanism for interacting with the Ethereum blockchain via RPC calls. It allows for the retrieval of transaction details using transaction hashes. The package includes interfaces and implementations necessary to fetch and process Ethereum transaction data.

#### Repositories

The repository package provides a data access layer for interacting with the `Transaction` and `User` models using the GORM library. This package abstracts the database operations related to transactions and users, offering a clean API for creating, updating, deleting, and retrieving transactions and users.

#### Models

The model package defines the schema for the `Transaction` and `User` models used with the GORM ORM library. This struct maps to a table in the database and represents a transaction or user with various attributes.

### Contract Package Overview

The contract package serves as a plugin for handling contract interactions and data for the the Ethereum `Fetcher` and `Listener`. It provides generated
bindings for smart contracts and also necessary handlers, repositories and models in order to be compatible with the `Fetcher` REST architecture. The package is provided to the `Fetcher` and `Listener` componenents in the initial Docker container initialization and is thus reused by both. Could be further extended with providing ABIs for multiple contracts.

![Alt text](xdiagrams/fetcher_contract.png)

#### Handles

##### Person Handler

The `PersonHandler` provides associated methods for handling HTTP requests related to Person entities defined in the `SimplePersonInfoContract`, interacting with both a database and an Ethereum smart contract. Intended to be used by the Fetcher `Handler Manager`.

#### Repositories

The repository package provides a data access layer for interacting with the `Person` contract model using the GORM library. This package abstracts the database operations related to `Person`, offering a clean API for creating, updating, deleting, and retrieving a `Person` data entry.

#### Models

The model package defines the schema for the `Person` model used with the GORM ORM library. This struct maps to a table in the database and represents a `Person` entry with various attributes.

### Listener Overview

The `Listener` is used for listening for updates to information on the Ethereum blockchain and synchronizes these updates with a local PostgreSQL database. It subscribes to blockchain events, processes these events, and updates the local database accordingly.

#### Listeners

##### PIC Listener

Currently the `Listener` consists of a single listener for `SimplePersonInfoContract`. The `PIC Listener` subscribes to the `PersonInfoUpdated`(provided by the contract package) events from a deployed `SimplePersonInfoContract`. If the `Listener` manages to retrieve an event log and unpack proper data(conforming to the contract model definitions from the contract package) it saves it to the database.

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
