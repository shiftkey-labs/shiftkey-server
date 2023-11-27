# Shiftkey App API

This repository contains the API code for the Shiftkey App, which is built using Go and Gin framework. The API interacts with a PostgreSQL database to store and retrieve data.

## Prerequisites

Before running the API, make sure you have the following installed:

- Go (version X.X.X)
- Gin (version X.X.X)
- PostgreSQL (version X.X.X)

## Getting Started

To get started with the Shiftkey App API, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/shiftkeyapp/api.git
   ```

2. Install the dependencies:

   ```shell
   go mod download
   ```

3. Set up the PostgreSQL database:

   - Create a new database named `shiftkey`.
   - Update the database connection details in the `config.yaml` file.

4. Build and run the API:

   ```shell
   go run main.go
   ```

## API Endpoints

The Shiftkey App API provides the following endpoints:

- `/users`: Manage user data.
- `/products`: Manage product data.
- `/orders`: Manage order data.

For detailed information about each endpoint, refer to the API documentation.

## Contributing

If you would like to contribute to the Shiftkey App API, please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the [MIT License](LICENSE).
