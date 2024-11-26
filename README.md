# Golang API with JWT Authentication and MongoDB

This project is a simple API built with Golang, featuring JWT-based authentication and MongoDB as the database. The API provides endpoints for user authentication and managing items, with Swagger documentation included.

Features
User Authentication: Login and register functionality using JWT.
MongoDB Integration: Data is stored in a MongoDB database.
Swagger Documentation: Easily explore and test API endpoints.
Gin Framework: Lightweight and fast HTTP framework.
Modular Design: Clean separation of concerns for better maintainability.

# Installation

Prerequisites
Go (v1.18 or higher)
MongoDB (running locally or on a server)
Swagger CLI (for generating documentation)

Clone the Repository
```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository

### Install Dependencies
```bash
go mod tidy

# Configuration

### MongoDB Connection
### Update the MongoDB connection string in config/database.go:
```bash
client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

