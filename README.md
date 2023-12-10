# My Golang API Project

## Description

This project is a simple Golang API designed for learning purposes. It uses PostgreSQL as a database, sqlx for querying, dotenv for configuration management, and gin-gonic for routing and handling requests.

## Features

- PostgreSQL: Robust and powerful open-source relational database.
- SQLX: An extension of the standard database/sql library, providing additional functionality.
- Dotenv: A module for loading environment variables from a .env file.
- Gin-Gonic: A high-performance web framework for building APIs.

## Getting started

### Prerequisites

Ensure you have the following installed:

- Go programming language
- PostgreSQL database
- Any necessary environment variables set in a .env file

- Clone:

```bash
  git clone https://github.com/victorbonato/books-go.git && cd books-go
```

- Download modules:

```bash
  go mod tidy
```

### Running

1. Start the Database:

```bash
  make postgres
```

This command will initialize the PostgreSQL database.

2. Run Initial Migrations:

```bash
  make migrate
```

This step involves setting up the database schema as per the migrations defined.

3. Run the Server:

```bash
  make run
```

This will start the Golang API server, making it ready to handle requests.

### Usage

Once the server is running, you can interact with the API through the defined endpoints. Use tools like Postman or curl for testing and interacting with your API.
Contributing

Feel free to fork the project and submit pull requests. Any contributions you make are greatly appreciated.

### License

MIT
