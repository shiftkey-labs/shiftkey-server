# Shiftkey App Server

This repository contains the API code for the Shiftkey App, which is built using Go and Gin framework. The API interacts with a PostgreSQL database to store and retrieve data.

## Index

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
  - [Users](#users)
  - [Events](#events)
  - [Attendance](#attendance)
  - [Hosts](#hosts)
  - [Authentication](#authentication)
- [Database Schema](#database-schema)
- [Contributing](#contributing)
- [License](#license)

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

### Users

- **GET `/users`**: Retrieve a list of users.
- **POST `/users`**: Create a new user.
- **GET `/users/:id`**: Retrieve a specific user by ID.
- **PUT `/users/:id`**: Update a specific user by ID.
- **DELETE `/users/:id`**: Delete a specific user by ID.

### Events

- **GET `/events`**: Retrieve a list of events.
- **POST `/events`**: Create a new event.
- **GET `/events/:id`**: Retrieve a specific event by ID.
- **PUT `/events/:id`**: Update a specific event by ID.
- **DELETE `/events/:id`**: Delete a specific event by ID.

### Attendance

- **GET `/attendance`**: Retrieve attendance records.
- **POST `/attendance`**: Create an attendance record.
- **PUT `/attendance/:id`**: Update an attendance record by ID.
- **DELETE `/attendance/:id`**: Delete an attendance record by ID.

### Hosts

- **GET `/hosts`**: Retrieve a list of hosts.
- **POST `/hosts`**: Create a new host.
- **GET `/hosts/:id`**: Retrieve details of a specific host.
- **PUT `/hosts/:id`**: Update details of a specific host.
- **DELETE `/hosts/:id`**: Delete a specific host.

### Authentication

- **POST `/auth/login`**: Authenticate a user and return a token.
- **POST `/auth/register`**: Register a new user.
- **POST `/auth/logout`**: Log out a user.

## Database Schema

![Database Schema](https://res.cloudinary.com/dhaaujyea/image/upload/v1703309710/sk_server_x8uoyh.png)

## Contributing

If you would like to contribute to the Shiftkey App API, please follow the guidelines in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the [MIT License](LICENSE).
