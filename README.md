# apm-server-elastic-go

This is a sample Go application with a well-structured project layout using the Gin framework, GORM for database interactions, and Elastic APM for monitoring.

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Elastic APM Integration](#elastic-apm-integration)

## Project Structure

```
myapp/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
│   └── config.json
├── controllers/
│   └── user_controller.go
├── models/
│   └── user.go
├── repository/
│   └── user_repository.go
├── routes/
│   └── routes.go
├── middleware/
│   └── elastic.go
├── go.mod
├── go.sum
└── README.md
```

## Installation

1. Clone the repository:
```sh
git clone https://github.com/balasl342/apm-server-elastic-go.git
```

2. Install dependencies:
```sh
go mod tidy
```

3. Set up the database and update `config.json` with your database credentials.

## Configuration

The configuration file `config/config.json` should contain the below application settings:
```json
{
    "database": {
        "DBHost": "your_db_host",
        "DBPort": "your_db_port",
        "DBUser": "your_db_user",
        "DBPassword": "your_db_password",
        "DBName": "your_db_name"
    },
    "elastic_apm": {
        "APMServerURLs": "http://localhost:8200",
        "APMServiceName": "your_apm_service_name",
        "APMSecretToken": "your_apm_secret_token",
        "APMEnvironment": "your_apm_environment"
    }
}
```

## Running the Application

1. Run the application:
```sh
go run cmd/server/main.go
```

2. The server will start on port 8080. You can change the port by updating `main.go`.

## API Endpoints

### Get all users
```http
GET /get_all_users
```

### Get user by ID
```http
GET /get_user/{id}
```

### Create a new user
```http
POST /create_user
{
    "name": "John Doe",
    "email": "john.doe@example.com"
}
```

## Elastic APM Integration

Elastic APM is integrated to monitor the performance of the application. The configuration is handled in `middleware/elastic.go`.

Update your `config.json` with the Elastic APM server URL and service name.

### Visualize in Elastic APM

- To view the data collected by elastic:
    * Log in to your elastic cloud account.
    * Navigate to the Observability section and choose APM.
    * Select your service from the list.
    * Explore the transactions, dependencies, metrics and logs that was created using this application.

## License

This project is licensed under the MIT License.
