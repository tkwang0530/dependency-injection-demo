# Dependency Injection Demo

This project is a simple TODO list RESTful API using Go, Gin, and Uber's `dig` for dependency injection. The project is designed to demonstrate modular and maintainable code using interfaces.

## How to Start

1. **Clone the repository**

    ```bash
    git clone https://github.com/tkwang0530/dependency-injection-demo.git
    cd dependency-injection-demo
    ```

2. **Initialize the Go module and install dependencies**

    ```bash
    go mod tidy
    ```

3. **Run the application**

    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:8080`.

## Folder Structure
- **main.go**: Entry point
- **di/**: Dependency Injection setup
  - **manager.go**: Centralized DI configuration
  - **config.go**: Config dependencies
  - **db.go**: Database dependencies
  - **repository.go**: Repository dependencies
  - **service.go**: Service dependencies
  - **handler.go**: Handler dependencies
- **config/**: Configuration
  - **config.go**: Configuration settings
- **db/**: Database layer
  - **in_memory_db.go**: Mock in-memory database
- **models/**: Data models
  - **todo.go**: TODO model
- **repository/**: Repository layer
  - **todo_repository.go**: TODO repository implementation
- **service/**: Service layer
  - **todo_service.go**: TODO service implementation
- **handler/**: HTTP handlers
  - **todo_handler.go**: TODO HTTP handler
- **router/**: Router setup
  - **router.go**: Gin router setup


## Features

- **Dependency Injection**: Uses Uber's `dig` for managing dependencies, making the code modular and maintainable.
- **RESTful API**: Provides a simple TODO list API with the following endpoints:
  - `POST /todos`: Add a new TODO item
  - `GET /todos`: Retrieve all TODO items
  - `DELETE /todos/:id`: Delete a TODO item by ID
- **In-Memory Database**: Uses an in-memory database to store TODO items for simplicity.
- **Gin Framework**: Utilizes the Gin web framework for handling HTTP requests and routing.

## Example API Usage

- **Add a new TODO item**

    ```bash
    curl -X POST http://localhost:8080/todos -d '{"task": "Buy milk"}' -H "Content-Type: application/json"
    ```

- **Retrieve all TODO items**

    ```bash
    curl http://localhost:8080/todos
    ```

- **Delete a TODO item by ID**

    ```bash
    curl -X DELETE http://localhost:8080/todos/1
    ```

Feel free to contribute to this project by creating pull requests or opening issues. Enjoy coding!