# Gin REST API

RESTful API with JWT authentication, role-based access, and structured routing using Gin.

## Skills
Gin, JWT, middleware, REST, structured routing

## Quick Start
```bash
go mod tidy
go run main.go
```
Server: http://localhost:8080

## Features
- User registration and login with JWT  
- Protected routes requiring valid JWT tokens  
- Role-based access control (user vs admin)  
- Request validation and error handling  
- Structured project layout: handlers, middleware, models, utils  
- In-memory database (no external dependencies)  
- Pre-seeded test users  

## Endpoints
Public  
- POST /api/register – Register new user  
- POST /api/login – Login and get JWT token  

Protected (require JWT)  
- GET  /api/profile – Get current user profile  
- PUT  /api/profile – Update user profile  
- GET  /api/users – List all users  

Admin (require admin role)  
- GET /api/admin/stats – Get system statistics  

## Test Users
- Admin: admin@example.com / admin123  
- User: user@example.com / user123

## Dependencies
- github.com/gin-gonic/gin  
- github.com/golang-jwt/jwt/v5  
- golang.org/x/crypto/bcrypt

MIT License