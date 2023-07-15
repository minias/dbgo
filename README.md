# DBGo - Dynamic Database Builder

DBGo is a dynamic builder that allows you to create database tables dynamically based on JSON input. It is built with Go and utilizes the Echo framework for handling HTTP requests.

**Note: This Project was generated using ChatGPT, a language model developed by OpenAI.**

## Prerequisites

Before running the application, make sure you have the following prerequisites installed:

- Go (version 1.16 or higher)
- MySQL or Oracle database (based on your choice)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/minias/dbgo.git
```

1. install the required dependencies:

```sh
go mod download
```

## Configuration

The application's configuration is stored in the config.yml file. Update the file with your desired settings:

- Server.Port: The port number on which the server will run.
- Server.Protocol: The protocol used by the server. Can be "http" or "https".
- Server.CertFile: Path to the SSL certificate file required for HTTPS protocol.
- Server.KeyFile: Path to the private key file required for HTTPS protocol.
- JWTSecret: Secret key used for JWT token signing.
- JWTExpireTime: Expiration time (in minutes) for JWT tokens.
- Database.Driver: Name of the database driver. Choose either "mysql" or "oracle".
- Database.Source: Database connection string.

## Usage

1. Start the application:

```sh
go run main.go
```

1. Access the API endpoints using a REST client or web browser:

- To create a table dynamically, send a PUT request to /table/create with the JSON data representing the table structure.
- To alter an existing table, send a PUT request to /table/alter with the JSON data representing the alterations.
- To delete a table, send a DELETE request to /table/delete with the table name as a parameter.
- To sign up as a user, send a POST request to /user/signup with the user email and password.
- To sign out, send a GET request to /user/signout.
- To log in, send a POST request to /user/login with the user email and password.
- To create a new database, send a PUT request to /database/create with the database name.
- To delete a database, send a DELETE request to /database/delete with the database name.

## License

This project is licensed under the [CC0 1.0 Universal](./LICENSE).


