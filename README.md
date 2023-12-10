# My Golang API Project

## Description
This project is a simple Golang API designed for learning purposes. It utilizes PostgreSQL for the database, sqlx for querying, dotenv for configuration management, and gin-gonic for efficient routing and handling.

## Features
- **PostgreSQL**: A robust and powerful open-source relational database.
- **SQLX**: Provides additional functionality to Go's standard database/sql library.
- **Dotenv**: Loads environment variables from a .env file.
- **Gin-Gonic**: A high-performance web framework for building APIs.

## Getting Started
### Prerequisites
- Go programming language
- PostgreSQL database
- `.env` file with necessary environment variables

### Installation
1. **Clone the Repository**:
    ```bash
    git clone https://github.com/victorbonato/books-go.git && cd books-go
    ```

2. **Download Modules**:
    ```bash
    go mod tidy
    ```

### Running the API
1. **Start the Database**:
    ```bash
    make postgres
    ```
2. **Run Initial Migrations**:
    ```bash
    make migrate
    ```
3. **Run the Server**:
    ```bash
    make run
    ```

### Usage
Once the server is running, you can interact with the API through the defined endpoints. Use tools like Postman or curl for testing and interacting with your API. Contributing

## Contributing
Feel free to fork the project and submit pull requests. Any contributions you make are greatly appreciated.


## License
[MIT](https://github.com/victorbonato/books-go/blob/main/LICENSE)
