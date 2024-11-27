# Golang API with JWT Authentication and MongoDB

This project is a simple API built with Golang, featuring JWT-based authentication and MongoDB as the database. The API provides endpoints for user authentication and managing items, with Swagger documentation included.

## Features

1. User Authentication: Login and register functionality using JWT.
2. MongoDB Integration: Data is stored in a MongoDB database.
3. Swagger Documentation: Easily explore and test API endpoints.
4. Gin Framework: Lightweight and fast HTTP framework.
5. Modular Design: Clean separation of concerns for better maintainability.

## Installation

### Prerequisites

Go (v1.23.3)
MongoDB (running locally or on a server)
Swagger CLI (for generating documentation)

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

| Method | Endpoint | Description |
| POST | | /register | Register a new user |
| POST | /login | Login and get a JWT |

#### Items

| Method | Endpoint | Description |
| GET | /items | Get a list of all items |
| GET | /items/{id} | Get details of a specific item |
| POST | /items | Add new item |
