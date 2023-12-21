# Shiftkey App API

This repository contains the API code for the Shiftkey App, which is built using Go and Gin framework. The API interacts with a PostgreSQL database to store and retrieve data.

## Prerequisites

Before running the API, make sure you have the following installed:

- Go
- PostgreSQL

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

   - Create a new database named `shiftkey-db`.
   - Create a new `.env` file using the `.env.example` file as a template.

4. Build and run the API:

   ```shell
   go run cmd/main/main.go
   ```

   or build and run the API using `run.sh`:

   ```shell
   ./run.sh
   ```

   Make sure to give permission to the script before running it:

   ```shell
   chmod +x run.sh
   ```

## API Endpoints

The Shiftkey App API provides the following endpoints:

- `/users`: Manage user data.
- `/events`: Manage events data.
- `/attendance`: Manage user attendance in events data.

For detailed information about each endpoint, refer to the API documentation.

## Contributing

If you would like to contribute to the Shiftkey App API, please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the [MIT License](LICENSE).
