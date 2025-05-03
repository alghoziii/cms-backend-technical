# CMS Backend API with Golang

REST API for Content Management System using Gin, GORM, PostgreSQL, JWT Auth, and Swagger.

## Features
- Auth with JWT
- CRUD: Categories, News, Custom Pages, Comment

## Teknologi 

- **Golang**
- **GORM**
- **PostgreSQL**
- **Swagger UI for documentation**

  ## Instalasi

1. Clone the repository:

   ```bash
   https://github.com/alghoziii/cms-backend-technical/

2. Install dependencies:
   ```bash
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    go get -u github.com/golang-jwt/jwt/v4
    go get -u github.com/joho/godotenv
    go get -u golang.org/x/crypto
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/files
3. Create .env file:
   ```bash
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=cms_db
    DB_PORT=5432
4. Create table products and orders:
   ```bash
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP
    );
    
    CREATE TABLE categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP
    );
    
    CREATE TABLE news (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE,
        user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP
    );
    
    CREATE TABLE custom_pages (
        id SERIAL PRIMARY KEY,
        url VARCHAR(255) UNIQUE NOT NULL,
        content TEXT NOT NULL,
        user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP
    );
    
    CREATE TABLE comments (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        comments TEXT NOT NULL,
        news_id INTEGER REFERENCES news(id) ON DELETE CASCADE,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP
    );
   
5. Running the Application:
   ```bash
   go run cmd/main.go

   
## Link Penting :
1. API Dokumentasi :
   
    - Swagger : http://localhost:8080/docs/index.html
      
3. API Aplikasi : http://localhost:8080/
  
4.  ![ERD Diagram](https://drive.google.com/file/d/1FwyRbKIbg3ctGjTeijTPNRhCjfjhhleV/view?usp=sharing)



   





