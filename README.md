# MySQL API with Go (Golang)

It contains a simple Go (Golang) application that serves as an API for managing accounts in a MySQL database. The application uses the Gin framework for routing and GORM for database interactions. Below are optimizations and explanations for the code components:

## `config` Package

### `config.Connect(dsn string) *gorm.DB`

The `Connect` function establishes a connection to a MySQL database using the provided Data Source Name (DSN). It also performs auto-migration to ensure that the database schema matches the defined Go data models. Error handling is in place for a more robust application. The function returns the GORM database object for further use.

## `controller` Package

The `controller` package contains functions for basic CRUD operations on the "Accounts" model in the database.

### `GetAcc(ctx *gin.Context)`

- Handles GET requests to retrieve account data.
- Uses GORM to fetch account records from the database.
- Proper error handling is implemented.

### `CreateAcc(ctx *gin.Context)`

- Handles POST requests to create a new account.
- Binds JSON data to the "Accounts" model and creates a new record in the database.
- Proper error handling is implemented.

### `UpdateAcc(ctx *gin.Context)`

- Handles PUT requests to update an existing account.
- Retrieves the account by UID, updates its attributes, and saves it back to the database.
- Proper error handling is implemented.

### `DelAcc(ctx *gin.Context)`

- Handles DELETE requests to delete an account by UID.
- Retrieves the account by UID and deletes it from the database.
- Proper error handling is implemented.

## `models` Package

The `models` package defines the data model for the "Account" table in the database using GORM tags. The code has been optimized for consistency and follows Go naming conventions.

## `routes` Package

### `routes.AccRoute(router *gin.Engine)`

- Defines routes for account management by grouping them under "/Accounts."
- Optimizes endpoint definitions for readability and maintenance.

## `main` Package

The entry point of the application initializes the database connection, sets up the Gin router, defines routes, and starts the server. Error handling is implemented when starting the server.

## How to Run

1. Update the MySQL DSN in `main.go` with your database connection details.
   or
   To create new database [connection with docker](dokcer/README.md)

2. Install the necessary Go dependencies.

   ```shell
   go mod tidy

3. Build and run the application.
    ```shell
   go run main.go
The server will start and listen on port 3000.


## Author

- Rajat Kumar Goyal

## License
This project is licensed under the [MIT LICENSE](LICENSE.md).
