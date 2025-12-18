# User Management API (Go + MySQL)

A RESTful API built with **GoFiber** to manage users with their name and date of birth.  
The API dynamically calculates a user’s age when fetching user details.

---

## 🚀 Tech Stack

- **Go** (Backend)
- **GoFiber** – Web framework
- **MySQL** – Database
- **SQLC** – Type-safe database access
- **Uber Zap** – Structured logging
- **go-playground/validator** – Request validation

---

## 📁 Project Structure

.
├── cmd/server/main.go
├── config/
├── db/
│   ├── migrations/
│   └── sqlc/
├── internal/
│   ├── handler/
│   ├── repository/
│   ├── service/
│   ├── routes/
│   ├── middleware/
│   ├── models/
│   └── logger/
├── go.mod
├── go.sum
├── sqlc.yaml
└── README.md



---

## 🗄️ Database Schema

```sql
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);
```
---

## ⚙️ Setup Instructions

1. **Prerequisites**

   - Go 1.21+
   - MySQL 8.x
   - MySQL Workbench (optional)
   - SQLC installed

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

2. **Database Setup**

  - Create database:

    ```sql
    CREATE DATABASE users_db;
    ```

  - Create application user:
    ```sql
    CREATE USER 'go_user'@'localhost'
    IDENTIFIED WITH mysql_native_password
    BY 'go_pass123';
    
    GRANT ALL PRIVILEGES ON users_db.* TO 'go_user'@'localhost';
    FLUSH PRIVILEGES;
    ```
    
3. **Generate SQLC Code**

    sqlc generate

4. **Run the Application**

    go run cmd/server/main.go

  - Server starts on:

    http://localhost:8080


📌 **API Endpoints**

    - Create User

    - POST /users
```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```
    - Response:

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}

**Get User by ID (with Age)**

**GET** /users/{id}

    - Response:

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

**Update User**

**PUT** /users/{id}

{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

**Delete User**

**DELETE** /users/{id}

    - Response:

    204 No Content

**List Users**

**GET** /users

    - Response (empty):

      []

    - Response (with data):

[
  {
    "id": 2,
    "name": "Bob",
    "dob": "1995-07-12",
    "age": 29
  }
]

---

**🧠 Design Notes**

- dob is stored in the database; **age is calculated dynamically** using Go’s time package

- SQLC enforces strict type safety between SQL and Go

- Clean architecture:

    - Handler -> Service -> Repository -> SQLC

- API returns empty arrays ([]) instead of null for better client compatibility.
