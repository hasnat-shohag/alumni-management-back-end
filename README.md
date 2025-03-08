# Alumni Management System

The Alumni Management System is a web application designed to manage alumni information, events, blogs, and executive committee members for an educational institution. The system provides functionalities for user authentication, profile management, event management, blog management, and more.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Features

- User authentication and authorization
- Profile management for alumni and students
- Event creation, updating, and deletion
- Blog creation and management
- Executive committee management
- Email notifications for user verification and password reset

## Frontend Codebase

You can find the frontend codebase for the Alumni Management System at the following link: [Alumni Management System Frontend](https://github.com/touhid9teen/Alumni-Management-System-ICE-RU)

### Key Directories and Files

- [`pkg/common`](pkg/common): Contains common utilities like logging and response handling.
- [`pkg/config`](pkg/config): Configuration management using environment variables.
- [`pkg/connection`](pkg/connection): Database connection setup.
- [`pkg/containers`](pkg/containers): Application container setup and server initialization.
- [`pkg/controllers`](pkg/controllers): Controllers for handling HTTP requests.
- [`pkg/domain`](pkg/domain): Domain interfaces for repositories and services.
- [`pkg/email`](pkg/email): Email templates and sending functionality.
- [`pkg/middlewares`](pkg/middlewares): Middleware for request validation and authentication.
- [`pkg/models`](pkg/models): Data models for the application.
- [`pkg/repositories`](pkg/repositories): Database interaction logic.
- [`pkg/routes`](pkg/routes): API route definitions.
- [`pkg/serializer`](pkg/serializer): Request and response serializers and validators.
- [`pkg/services`](pkg/services): Business logic and service layer.
- [`images`](images): Directory for storing uploaded images.
- [`main.go`](main.go): Entry point of the application.
- [`go.mod`](go.mod) and [`go.sum`](go.sum): Go module dependencies.
- [`app.env`](app.env) and [`base.env`](base.env): Environment variable files.

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/alumni-management-system.git
cd alumni-management-system
```

2. Install dependencies:

```sh
go mod tidy
```

3. Set up the environment variables:

Create a `.env` file in the root directory and add the necessary environment variables. You can use `app.env` and `base.env` as references.

4. Run the application:

```sh
go run main.go
```

## Usage

The application provides a RESTful API for managing alumni, events, blogs, and executive committee members. You can use tools like Postman or cURL to interact with the API.

## API Endpoints

### Authentication

- `POST /v1/auth/sign-up`: Sign up a new user.
- `POST /v1/auth/login`: Log in a user.

### User

- `POST /user/forget-password`: Request a password reset.
- `POST /user/reset-password`: Reset the password.
- `GET /user/ping`: Ping the server.
- `GET /user/alumni-list`: Get a list of all alumni.
- `GET /user/:id`: Get details of a specific alumni.
- `DELETE /user/delete-me/:id`: Delete the user's account.
- `PATCH /user/complete-profile/:id`: Update the user's profile.

### Admin

- `POST /v1/admin/verify-user`: Verify a user.
- `DELETE /v1/admin/delete-user/:id`: Delete a user.

### Event

- `POST /v1/event/create`: Create a new event.
- `PATCH /v1/event/update/:id`: Update an event.
- `DELETE /v1/event/delete/:id`: Delete an event.
- `GET /v1/event/:id`: Get details of a specific event.
- `GET /v1/events`: Get a list of all events.

### Blog

- `POST /v1/blog/create`: Create a new blog post.

### Executive Committee

- `POST /v1/executive-committee/create`: Add a new executive committee member.
- `PATCH /v1/executive-committee/update/:id`: Update an executive committee member.
- `DELETE /v1/executive-committee/delete/:id`: Delete an executive committee member.
- `GET /v1/executive-committee/list`: Get a list of all executive committee members.
- `GET /v1/executive-committee/:id`: Get details of a specific executive committee member.

## Environment Variables

The application uses environment variables for configuration. Here are the key variables:

- `DBUSER`: Database username
- `DBPASS`: Database password
- `DBIP`: Database IP address and port
- `DBNAME`: Database name
- `PORT`: Application port
- `JWT_SECRET`: Secret key for JWT
- `JWT_EXPIRE_MINUTES`: JWT expiration time in minutes
- `SMTP_HOST`: SMTP server host
- `SMTP_PORT`: SMTP server port
- `SMTP_USERNAME`: SMTP server username
- `SMTP_PASSWORD`: SMTP server password
- `ADMIN_EMAIL`: Admin email address

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
