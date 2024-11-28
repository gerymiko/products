# Golang API with JWT Authentication and MongoDB

This project is a simple API built with Golang, featuring JWT-based authentication and MongoDB as the database. The API provides endpoints for user authentication and managing items, with Swagger documentation included.

## Features

- User Authentication: Login and register functionality using JWT.
- MongoDB Integration: Data is stored in a MongoDB database.
- Swagger Documentation: Easily explore and test API endpoints.
- Gin Framework: Lightweight and fast HTTP framework.
- Modular Design: Clean separation of concerns for better maintainability.

## Installation

### Prerequisites

- Go (v1.23.3)
- MongoDB (running locally or on a server)
- Swagger CLI (for generating documentation)

### Clone the Repository

```bash
git clone https://github.com/gerymiko/products.git
cd your-repository
```

### Install Dependencies

```bash
go mod tidy
```

---

## Configuration

### MongoDB Connection

Update the MongoDB connection string in `config/database.go`:
Add to `.env` file

```bash
MONGO_URI=mongodb://example:3306
MONGO_DB_NAME=example-db
```

### Environment Variables

Create a `.env` file for sensitive configurations like JWT secrets:

```bash
JWT_SECRET=your-secret-key
```

### Usage

Run the Application
Start the API server:

```bash
go run main.go
```

### Swagger Documentation

Once the server is running, Swagger UI is available at:

```bash
http://localhost:8080/swagger/index.html
```

### API Endpoints

#### Authentication

| Method | Endpoint  | Description         |
| ------ | --------- | ------------------- |
| POST   | /register | Register a new user |
| POST   | /login    | Login and get a JWT |

#### Items

| Method | Endpoint    | Description                    |
| ------ | ----------- | ------------------------------ |
| GET    | /items      | Get a list of all items        |
| GET    | /items/{id} | Get details of a specific item |
| POST   | /items      | Add new item                   |

## Development

### Generate Swagger Documentation

After making changes to the handlers, regenerate the Swagger documentation:

```bash
swag init
```

### Project Structure

```bash
/products
  ├── main.go                # Entry point of the application
  ├── config/
  │   └── database.go        # MongoDB configuration
  ├── docs/                  # Auto-generated Swagger files
  ├── handlers/
  │   ├── auth.go            # Authentication handlers
  │   ├── items.go           # Item-related handlers
  ├── models/
  │   ├── user.go            # User model
  │   ├── item.go            # Item model
  ├── middlewares/
  │   └── auth.go            # JWT middleware
  ├── .env                   # Environment variables
  └── go.mod                 # Go modules
```

## Testing

Use tools like Postman or cURL to test the endpoints. For example:

### Register a New User

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username":"testuser","password":"testpass"}'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $TOKEN" \
-d '{"username":"testuser","password":"testpass"}'
```
To run unit test just run this code:
```bash
go test ./...
```