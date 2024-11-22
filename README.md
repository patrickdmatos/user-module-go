#User Module API

###Overview
The User Module API is a RESTful service built in Go to manage user-related operations such as registration, authentication, and data manipulation. It leverages the Fiber web framework for handling HTTP requests, GORM for interacting with PostgreSQL, and JWT for user authentication.

###Tools and Technologies
- Go 1.19: The programming language used to build the service.
- Fiber v2.52.5: A fast web framework for Go, used to handle HTTP requests and routing.
- GORM v1.25.12: ORM (Object-Relational Mapping) library for interacting with PostgreSQL and managing database operations.
- PostgreSQL: The database used to store user data.
- JWT v3.2.2: JSON Web Tokens for secure user authentication.
- godotenv v1.5.1: Loads environment variables from a .env file for configuration.

###Dependencies
This project uses the following dependencies:

- Fiber for web server and routing.
- GORM for database interaction (PostgreSQL driver).
- JWT for secure authentication.
- godotenv to manage environment variables.
- pgx and other indirect dependencies for PostgreSQL connection pooling and utilities.

###Setup
Clone the repository:

git clone https://github.com/yourusername/user-module-go.git
cd user-module-go
Install dependencies:

go mod tidy
Set up the database connection in the .env file:

Run the application:

bash
Copy code
go run cmd/main.go
The application will be available at http://localhost:3000.
